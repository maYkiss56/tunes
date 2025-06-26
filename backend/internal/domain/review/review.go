package review

import "time"

type Review struct {
	ID        int
	UserID    int
	SongID    int
	Body      string
	IsLike    bool
	IsValid   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewReview(userID int, songID int, body string, isLike bool, isValid bool) (*Review, error) {
	return &Review{
		UserID:    userID,
		SongID:    songID,
		Body:      body,
		IsLike:    isLike,
		IsValid:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
