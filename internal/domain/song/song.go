package song

import (
	"time"
)

type Song struct {
	ID          int
	Title       string
	FullTitle   string
	ImageURL    string
	ReleaseDate time.Time
	ArtistID    int
	AlbumID     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
