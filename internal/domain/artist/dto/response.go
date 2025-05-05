package dto

import "github.com/maYkiss56/tunes/internal/domain/artist"

type Response struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	BIO      string `json:"bio"`
	Country  string `json:"country"`
}

func ToResponse(a artist.Artist) Response {
	return Response{
		ID:       a.ID,
		Nickname: a.Nickname,
		BIO:      a.BIO,
		Country:  a.Country,
	}
}
