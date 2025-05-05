package dto

import "github.com/maYkiss56/tunes/internal/domain/album"

type Response struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	ArtistID int    `json:"artist_id"`
}

func ToResponse(a album.Album) Response {
	return Response{
		ID:       a.ID,
		Title:    a.Title,
		ImageURL: a.ImageURL,
		ArtistID: a.ArtistID,
	}
}
