package session

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	domain "github.com/maYkiss56/tunes/internal/domain/users"
)

type Session struct {
	ID         string
	UserID     int
	UserEmail  string
	UserRoleID int
	CreatedAt  time.Time
	ExpiresAt  time.Time
	Data       map[string]interface{}
}

func GenerateSession(r *http.Request, u *domain.User, rememberMe bool) (Session, error) {
	id := uuid.NewString()
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}

	expiry := time.Now().Add(24 * time.Hour)
	if rememberMe {
		expiry = time.Now().Add(30 * 24 * time.Hour)
	}

	session := Session{
		ID:         id,
		UserID:     u.ID,
		UserEmail:  u.Email,
		UserRoleID: u.RoleID,
		CreatedAt:  time.Now(),
		ExpiresAt:  expiry,
		Data: map[string]interface{}{
			"user_agent": r.UserAgent(),
			"ip":         ip,
		},
	}

	return session, nil
}
