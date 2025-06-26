package dto

import (
	"github.com/maYkiss56/tunes/internal/domain/album"
	"github.com/maYkiss56/tunes/internal/domain/artist"
	artistDTO "github.com/maYkiss56/tunes/internal/domain/artist/dto"
)

type Response struct {
	ID       int                `json:"id"`
	Title    string             `json:"title"`
	ImageURL string             `json:"image_url"`
	Artist   artistDTO.Response `json:"artist"`
}

func ToResponse(a album.Album, ar artist.Artist) Response {
	return Response{
		ID:       a.ID,
		Title:    a.Title,
		ImageURL: a.ImageURL,
		Artist: artistDTO.Response{
			ID:       ar.ID,
			Nickname: ar.Nickname,
			BIO:      ar.BIO,
			Country:  ar.Country,
		},
	}
}
