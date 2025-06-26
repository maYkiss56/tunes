package utilites

import (
	"net/http"
	"time"

	"github.com/maYkiss56/tunes/internal/session"
)

func SetCookie(w http.ResponseWriter, s session.Session) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    s.ID,
		Path:     "/",
		Expires:  s.ExpiresAt,
		MaxAge:   int(time.Until(s.ExpiresAt).Seconds()),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func CleanCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}
