package dto

import "errors"

type CreateAlbumRequest struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	ArtistID int    `json:"artist_id"`
}

func (r *CreateAlbumRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.ImageURL == "" {
		return errors.New("image_url is required")
	}
	if r.ArtistID == 0 {
		return errors.New("artist_id is required")
	}
	return nil

}

type UpdateAlbumRequest struct {
	Title    *string `json:"title,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
	ArtistID *int    `json:"artist_id,omitempty"`
}

func (r *UpdateAlbumRequest) Validate() error {
	if r.Title != nil && *r.Title == "" {
		return errors.New("title cannot be empty")
	}
	if r.ImageURL != nil && *r.ImageURL == "" {
		return errors.New("image_url cannot be empty")
	}
	if r.ArtistID != nil && *r.ArtistID == 0 {
		return errors.New("artist_id cannot be empty")
	}

	return nil
}
