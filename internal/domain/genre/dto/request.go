package dto

import "errors"

type CreateGenreRequest struct {
	Title string `json:"title"`
}

func (r *CreateGenreRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

type UpdateGenreRequest struct {
	Title *string `json:"title,omitempty"`
}

func (r *UpdateGenreRequest) Validate() error {
	if r.Title != nil && len(*r.Title) > 150 {
		return errors.New("title is too long")
	}

	return nil
}
