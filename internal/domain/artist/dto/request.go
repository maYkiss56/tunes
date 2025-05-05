package dto

import "errors"

const (
	NicknameLenght = 150
)

var (
	ErrNicknameLenght = errors.New("nickname is too long")
)

type CreateArtistRequest struct {
	Nickname string `json:"nickname"`
	BIO      string `json:"bio"`
	Country  string `json:"country"`
}

func (r *CreateArtistRequest) Validate() error {
	if len(r.Nickname) > NicknameLenght {
		return ErrNicknameLenght
	}
	return nil
}

type UpdateArtistRequest struct {
	Nickname *string `json:"nickname,omitempty"`
	BIO      *string `json:"bio,omitempty"`
	Country  *string `json:"country,omitempty"`
}

func (r *UpdateArtistRequest) Validate() error {
	if r.Nickname != nil && len(*r.Nickname) > NicknameLenght {
		return ErrNicknameLenght
	}

	return nil
}
