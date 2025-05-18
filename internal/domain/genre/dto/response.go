package dto

import "github.com/maYkiss56/tunes/internal/domain/genre"

type Response struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func ToResponse(s genre.Genre) Response {
	return Response{
		ID:    s.ID,
		Title: s.Title,
	}
}
