package dto

import (
	"time"

	"github.com/maYkiss56/tunes/internal/domain/song"
)

type Response struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	FullTitle   string     `json:"full_title,omitempty"`
	ImageURL    string     `json:"image_url,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	ArtistID    int        `json:"artist_id,omitempty"`
	AlbumID     int        `json:"album_id,omitempty"`
}

func ToResponse(s song.Song) Response {
	return Response{
		ID:          s.ID,
		Title:       s.Title,
		FullTitle:   s.FullTitle,
		ImageURL:    s.ImageURL,
		ReleaseDate: &s.ReleaseDate,
		ArtistID:    s.ArtistID,
		AlbumID:     s.AlbumID,
	}
}
