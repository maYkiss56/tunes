package moderation

import "time"

type Status string

const (
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
	StatusPending  Status = "pending"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusApproved, StatusRejected, StatusPending:
		return true
	}
	return false
}

type Moderation struct {
	ID          int
	ReviewID    int
	ModeratorID int
	Status      Status
	Reason      string
	ModeratedAt time.Time
}
