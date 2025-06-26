package middleware

import (
	"net/http"
	"time"

	"github.com/maYkiss56/tunes/internal/session"
	"github.com/maYkiss56/tunes/internal/utilites"
)

const AdminRoleID = 4

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			utilites.RenderError(
				w,
				r,
				http.StatusUnauthorized,
				"unauthorized: session cookie not found",
			)
			return
		}

		sess, ok := session.GetSession(cookie.Value)
		if !ok || sess.ExpiresAt.Before(time.Now()) {
			utilites.RenderError(
				w,
				r,
				http.StatusUnauthorized,
				"unauthorized: session expired or invalid",
			)
			return
		}

		// сохраняем сессию в контексте для последующего использования
		ctx := r.Context()
		ctx = session.WithSession(ctx, sess)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := session.FromContext(r.Context())
		if s == nil {
			utilites.RenderError(w, r, http.StatusUnauthorized, "unauthorized: no session found")
			return
		}

		if s.UserRoleID != AdminRoleID {
			utilites.RenderError(w, r, http.StatusForbidden, "forbidden: insufficient permissions")
			return
		}

		next.ServeHTTP(w, r)
	})
}
