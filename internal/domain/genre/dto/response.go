package dto

import (
	"github.com/maYkiss56/tunes/internal/domain/genre"
)

type Response struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageURl string `json:"image_url"`
}

func ToResponse(s genre.Genre) Response {
	return Response{
		ID:       s.ID,
		Title:    s.Title,
		ImageURl: s.ImageURL,
	}
}
