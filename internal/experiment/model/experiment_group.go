package model

import "time"

type ExperimentGroup struct {
	ID                 int       `db:"id"`
	ExperimentId       int       `db:"experiment_id"`
	Name               string    `db:"name"`
	Slug               string    `db:"slug"`
	IsActive           bool      `db:"is_active"`
	CoveragePercentage float64   `db:"coverage_percentage"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`

	UserCount      int `db:"user_count"`
	OccupancyRatio float64
}
