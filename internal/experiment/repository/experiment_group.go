package repository

import (
	"ab-testing/internal/experiment/model"
	"ab-testing/pkg/db"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ExperimentGroupRepository struct {
	Db *sqlx.DB
}

func NewExperimentGroupRepository(db *sqlx.DB) ExperimentGroupRepository {
	return ExperimentGroupRepository{Db: db}
}

func (r ExperimentGroupRepository) GetByExperimentId(ctx context.Context, experimentId int, isActive bool) ([]model.ExperimentGroup, error) {
	var entities []model.ExperimentGroup
	query := fmt.Sprintf(
		"SELECT * FROM experiment_groups WHERE experiment_id = $1 AND is_active = $2",
	)

	err := r.Db.SelectContext(ctx, &entities, query, experimentId, isActive)

	return entities, db.HandleError(err)
}

func (r ExperimentGroupRepository) Find(ctx context.Context, id int) (model.ExperimentGroup, error) {
	entity := model.ExperimentGroup{}
	query := fmt.Sprintf(
		"SELECT * FROM experiment_groups WHERE id = $1 AND is_active=true LIMIT 1",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)

	return entity, db.HandleError(err)
}

func (r ExperimentGroupRepository) GetActiveGroupsWithUserCount(ctx context.Context, experimentId int) ([]model.ExperimentGroup, error) {
	var entities []model.ExperimentGroup
	query := fmt.Sprintf("SELECT eg.*, (SELECT COUNT(*) FROM experiment_users eu WHERE eu.group_id = eg.id) AS user_count FROM experiment_groups eg WHERE eg.experiment_id = $1 AND eg.is_active = true AND eg.coverage_percentage > 0")

	err := r.Db.SelectContext(ctx, &entities, query, experimentId)

	return entities, db.HandleError(err)
}
