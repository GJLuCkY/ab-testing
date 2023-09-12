package repository

import (
	"ab-testing/internal/experiment/model"
	"ab-testing/pkg/db"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ExperimentUserRepository struct {
	Db *sqlx.DB
}

func NewExperimentUserRepository(db *sqlx.DB) ExperimentUserRepository {
	return ExperimentUserRepository{Db: db}
}

func (r ExperimentUserRepository) GetByUserIdAndExperimentId(ctx context.Context, experimentId int, userID string) (model.ExperimentUser, error) {
	entity := model.ExperimentUser{}
	query := fmt.Sprintf(
		"SELECT * FROM experiment_users WHERE experiment_id = $1 AND user_id = $2 LIMIT 1",
	)
	err := r.Db.GetContext(ctx, &entity, query, experimentId, userID)

	return entity, db.HandleError(err)
}

func (r ExperimentUserRepository) Create(ctx context.Context, entity model.ExperimentUser) error {
	query := `INSERT INTO experiment_users (experiment_id, group_id, user_id)
				VALUES (:experiment_id, :group_id, :user_id);`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return db.HandleError(err)
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return db.HandleError(err)
		}
	}

	return db.HandleError(err)
}
