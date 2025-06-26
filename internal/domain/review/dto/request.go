package dto

import "errors"

type CreateReviewRequest struct {
	UserID  int    `json:"user_id"`
	SongID  int    `json:"song_id"`
	Body    string `json:"body"`
	IsLike  bool   `json:"is_like"`
	IsValid bool   `json:"is_valid"`
}

func (r *CreateReviewRequest) Validate() error {
	if r.UserID == 0 {
		return errors.New("user_id is required")
	}
	if r.SongID == 0 {
		return errors.New("song_id is required")
	}
	if r.Body == "" {
		return errors.New("text review is required")
	}
	return nil
}

type UpdateReviewRequest struct {
	Body   *string `json:"body,omitempty"`
	IsLike *bool   `json:"is_like,omitempty"`
	SongID int     `json:"song_id"`
}
