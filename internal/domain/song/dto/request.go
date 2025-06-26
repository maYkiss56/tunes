package dto

import (
	"errors"
	"time"
)

const (
	maxTitleLength     = 50
	maxFullTitleLength = 150
)

type CreateSongRequest struct {
	Title       string    `json:"title"`
	FullTitle   string    `json:"full_title"`
	ImageURL    string    `json:"image_url"`
	ReleaseDate time.Time `json:"release_date"`
	GenreID     int       `json:"genre_id"`
	ArtistID    int       `json:"artist_id"`
	AlbumID     int       `json:"album_id,omitempty"`
}

func (r *CreateSongRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if len(r.Title) > maxTitleLength {
		return errors.New("title is too long")
	}

	if r.FullTitle == "" {
		return errors.New("full title is required")
	}
	if len(r.FullTitle) > maxFullTitleLength {
		return errors.New("full title is too long")
	}

	if r.ImageURL == "" {
		return errors.New("image_url is required")
	}

	if r.ReleaseDate.IsZero() {
		return errors.New("release_date is required")
	}

	return nil
}

type UpdateSongRequest struct {
	Title       *string    `json:"title,omitempty"`
	FullTitle   *string    `json:"full_title,omitempty"`
	ImageURL    *string    `json:"image_url,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	GenreID     *int       `json:"genre_id,omitempty"`
	ArtistID    *int       `json:"artist_id,omitempty"`
	AlbumID     *int       `json:"album_id,omitempty"`
}

func (r *UpdateSongRequest) Validate() error {
	if r.Title != nil && len(*r.Title) > maxTitleLength {
		return errors.New("title is too long")
	}

	if r.FullTitle != nil && len(*r.FullTitle) > maxFullTitleLength {
		return errors.New("full title is too long")
	}

	return nil
}
