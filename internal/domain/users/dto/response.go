package dto

import "github.com/maYkiss56/tunes/internal/domain/users"

type Response struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	// AvatarURL string `json:"avatar_url,omitempty"`
	// IsBanned  bool   `json:"is_banned"`
	RoleID int `json:"role_id"`
}

func ToResponse(u users.User) Response {
	return Response{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
		// AvatarURL: u.AvatarURL,
		// IsBanned:  u.IsBanned,
		RoleID: u.RoleID,
	}
}
