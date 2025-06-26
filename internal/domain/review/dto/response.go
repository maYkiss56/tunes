package dto

import (
	"time"

	"github.com/maYkiss56/tunes/internal/domain/review"
	"github.com/maYkiss56/tunes/internal/domain/song"
	songDTO "github.com/maYkiss56/tunes/internal/domain/song/dto"
	"github.com/maYkiss56/tunes/internal/domain/users"
	userDTO "github.com/maYkiss56/tunes/internal/domain/users/dto"
)

type Response struct {
	ID        int              `json:"id"`
	User      userDTO.Response `json:"user"`
	Song      songDTO.Response `json:"song"`
	Body      string           `json:"body"`
	IsLike    bool             `json:"is_like"`
	IsValid   bool             `json:"is_valid"`
	UpdatedAt time.Time        `json:"updated_at"`
}

func ToResponse(r review.Review, u users.User, s song.Song) Response {
	return Response{
		ID: r.ID,
		User: userDTO.Response{
			ID:        u.ID,
			Email:     u.Email,
			Username:  u.Username,
			AvatarURL: u.AvatarURL,
		},
		Song: songDTO.Response{
			ID:          s.ID,
			Title:       s.Title,
			FullTitle:   s.FullTitle,
			ImageURL:    s.ImageURL,
			ReleaseDate: &s.ReleaseDate,
		},
		Body:      r.Body,
		IsLike:    r.IsLike,
		IsValid:   r.IsValid,
		UpdatedAt: r.UpdatedAt,
	}
}
