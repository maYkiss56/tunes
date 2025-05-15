package dto

import "errors"

type CreateUsersRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUsersRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Username == "" {
		return errors.New("username is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	if len(r.Password) < 8 {
		return errors.New("password must be more than 8 symbols")
	}

	return nil
}

type LoginUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	RemeberMe bool   `json:"remember_me"`
}

func (r *LoginUserRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	if len(r.Password) < 8 {
		return errors.New("password must be more than 8 symbols")
	}

	return nil
}

type UpdateAvatarRequest struct {
	AvatarURL string `json:"avatar_url"`
}

func (r *UpdateAvatarRequest) Validate() error {
	if r.AvatarURL == "" {
		return errors.New("avatar is required")
	}

	return nil
}

type UpdatePasswordRequest struct {
	NewPassword string `json:"new_password"`
}

func (r *UpdatePasswordRequest) Validate() error {
	if r.NewPassword == "" {
		return errors.New("new_password is required")
	}
	if len(r.NewPassword) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

type UpdateUsersRequest struct {
	Email     *string `json:"email,omitempty"`
	Username  *string `json:"username,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

func (r *UpdateUsersRequest) Validate() error {
	if r.Email != nil && len(*r.Email) > 150 {
		return errors.New("email is too long")
	}

	return nil
}
