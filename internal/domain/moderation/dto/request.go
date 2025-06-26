package dto

import (
	"errors"

	"github.com/maYkiss56/tunes/internal/domain/moderation"
)

type CreateModerationRequest struct {
	ReviewID    int    `json:"review_id"`
	ModeratorID int    `json:"moderator_id"`
	Status      string `json:"status"`
	Reason      string `json:"reason"`
}

func (r *CreateModerationRequest) Validate() error {
	if r.ReviewID == 0 {
		return errors.New("review_id is required")
	}
	if r.ModeratorID == 0 {
		return errors.New("moderator_id is required")
	}
	if r.Status == "" {
		return errors.New("status is required")
	}
	if !moderation.Status(r.Status).IsValid() {
		return errors.New("invalid status")
	}
	if r.Reason == "" {
		return errors.New("reason is required")
	}

	return nil
}

type UpdateModerationRequest struct {
	Status *string `json:"status"`
}

func (r *UpdateModerationRequest) Validate() error {
	if r.Status != nil && !moderation.Status(*r.Status).IsValid() {
		return errors.New("invalid status")
	}

	return nil
}
