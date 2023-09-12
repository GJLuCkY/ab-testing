package model

import "time"

type Status int

const (
	StatusPending Status = iota + 1
	StatusInProgress
	StatusDone
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending:
		return true
	case StatusInProgress:
		return true
	case StatusDone:
		return true
	}
	return false
}

type Experiment struct {
	ID             int       `db:"id"`
	Name           string    `db:"name"`
	Slug           string    `db:"slug"`
	Description    string    `db:"description"`
	IsActive       bool      `db:"is_active"`
	Stores         []byte    `db:"stores"`
	Platforms      []byte    `db:"platforms"`
	OnlyAuthorized bool      `db:"only_authorized"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
