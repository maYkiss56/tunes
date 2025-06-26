package song

import (
	"time"
)

type Song struct {
	ID           int
	Title        string
	FullTitle    string
	ImageURL     string
	ReleaseDate  time.Time
	LikeCount    int
	DislikeCount int
	Rating       int
	GenreID      int
	ArtistID     int
	AlbumID      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewSong(
	title string,
	fullTitle string,
	imageURL string,
	releaseData time.Time,
	genreID int,
	artistID int,
	albumID int,
) (*Song, error) {
	return &Song{
		Title:       title,
		FullTitle:   fullTitle,
		ImageURL:    imageURL,
		ReleaseDate: releaseData,
		GenreID:     genreID,
		ArtistID:    artistID,
		AlbumID:     albumID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
