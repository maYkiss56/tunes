package dto

import (
	"time"

	"github.com/maYkiss56/tunes/internal/domain/album"
	albumDTO "github.com/maYkiss56/tunes/internal/domain/album/dto"
	"github.com/maYkiss56/tunes/internal/domain/artist"
	artistDTO "github.com/maYkiss56/tunes/internal/domain/artist/dto"
	"github.com/maYkiss56/tunes/internal/domain/genre"
	genreDTO "github.com/maYkiss56/tunes/internal/domain/genre/dto"
	"github.com/maYkiss56/tunes/internal/domain/song"
)

type Response struct {
	ID           int                `json:"id"`
	Title        string             `json:"title"`
	FullTitle    string             `json:"full_title,omitempty"`
	ImageURL     string             `json:"image_url,omitempty"`
	ReleaseDate  *time.Time         `json:"release_date,omitempty"`
	LikeCount    int                `json:"like_count"`
	DislikeCount int                `json:"dislike_count"`
	Rating       int                `json:"rating"`
	Genre        genreDTO.Response  `json:"genre"`
	Artist       artistDTO.Response `json:"artist"`
	Album        albumDTO.Response  `json:"album"`
}

func ToResponse(s song.Song, g genre.Genre, songArtist artist.Artist, a album.Album, albumArtist artist.Artist) Response {
	return Response{
		ID:           s.ID,
		Title:        s.Title,
		FullTitle:    s.FullTitle,
		ImageURL:     s.ImageURL,
		ReleaseDate:  &s.ReleaseDate,
		LikeCount:    s.LikeCount,
		DislikeCount: s.DislikeCount,
		Rating:       s.Rating,
		Genre: genreDTO.Response{
			ID:       g.ID,
			Title:    g.Title,
			ImageURl: g.ImageURL,
		},
		Artist: artistDTO.Response{
			ID:       songArtist.ID,
			Nickname: songArtist.Nickname,
			BIO:      songArtist.BIO,
			Country:  songArtist.Country,
		},
		Album: albumDTO.Response{
			ID:       a.ID,
			Title:    a.Title,
			ImageURL: a.ImageURL,
			Artist: artistDTO.Response{
				ID:       albumArtist.ID,
				Nickname: albumArtist.Nickname,
				BIO:      albumArtist.BIO,
				Country:  albumArtist.Country,
			},
		},
	}
}
