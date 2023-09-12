package model

import "time"

type ExperimentUser struct {
	ID           int       `db:"id"`
	ExperimentID int       `db:"experiment_id"`
	GroupID      int       `db:"group_id"`
	UserID       int       `db:"user_id"`
	AnonymousID  string    `db:"anonymous_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
