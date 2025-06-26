package dto

import (
	"github.com/maYkiss56/tunes/internal/domain/moderation"
)

type Response struct {
	ID          int    `json:"id"`
	ReviewID    int    `json:"review_id"`
	ModeratorID int    `json:"moderator_id"`
	Status      string `json:"status"`
	Reason      string `json:"reason"`
}

func ToResponse(m moderation.Moderation) Response {
	return Response{
		ID:          m.ID,
		ReviewID:    m.ReviewID,
		ModeratorID: m.ModeratorID,
		Status:      string(m.Status),
		Reason:      m.Reason,
	}
}
