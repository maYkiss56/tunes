package users

import (
	"time"
)

const RoleID = 2 // user

type User struct {
	ID           int
	Email        string
	Username     string
	PasswordHash string
	AvatarURL    string
	IsBanned     bool
	RoleID       int
	LastLogin    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(email, username, password string) (*User, error) {
	return &User{
		Email:        email,
		Username:     username,
		PasswordHash: password,
		// AvatarURL:    "",
		// IsBanned:     false,
		RoleID:    RoleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) SetPasswordHash(hash string) {
	u.PasswordHash = hash
}
