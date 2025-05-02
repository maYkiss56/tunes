package song

import (
	"errors"
	"time"
)

const (
	maxTitleLength = 150
)

var (
	ErrNotFound = errors.New("song not found")
)

type Song struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	FullTitle   string    `json:"full_title"`
	ImageURL    string    `json:"image_url"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Filter struct {
	Title  string `json:"title"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type CreateSongRequest struct {
	Title       string    `json:"title"        validate:"required"`
	FullTitle   string    `json:"full_title"   validate:"required"`
	ImageURL    string    `json:"image_url"    validate:"required,url"`
	ReleaseDate time.Time `json:"release_date" validate:"required"`
}

func (r *CreateSongRequest) Validate() error {
	if len(r.Title) > maxTitleLength {
		return errors.New("title is too long")
	}
	return nil
}

type UpdateSongRequest struct {
	Title       *string    `json:"title,omitempty"`
	FullTitle   *string    `json:"full_title,omitempty"`
	ImageURL    *string    `json:"image_url,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
}

type Response struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	FullTitle   string     `json:"full_title,omitempty"`
	ImageURL    string     `json:"image_url,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
}

type ListResponse struct {
	Data   []Response `json:"data"`
	Total  int        `json:"total"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
}

func ToResponse(s Song) Response {
	return Response{
		ID:          s.ID,
		Title:       s.Title,
		FullTitle:   s.FullTitle,
		ImageURL:    s.ImageURL,
		ReleaseDate: &s.ReleaseDate,
	}
}

func ToListResponse(songs []Song, total, limit, offset int) ListResponse {
	data := make([]Response, 0, len(songs))
	for _, s := range songs {
		data = append(data, ToResponse(s))
	}

	return ListResponse{
		Data:   data,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}
}
