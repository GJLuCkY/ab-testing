package repository

import (
	"context"
	"fmt"

	"ab-testing/internal/experiment/model"
	"ab-testing/pkg/db"

	"github.com/jmoiron/sqlx"
)

type ExperimentRepository struct {
	Db *sqlx.DB
}

func NewExperimentRepository(db *sqlx.DB) ExperimentRepository {
	return ExperimentRepository{Db: db}
}

func (r ExperimentRepository) FirstBySlug(ctx context.Context, slug string, isActive bool) (model.Experiment, error) {
	entity := model.Experiment{}
	query := fmt.Sprintf(
		"SELECT * FROM experiments WHERE slug = $1 AND is_active = $2 LIMIT 1",
	)
	err := r.Db.GetContext(ctx, &entity, query, slug, isActive)

	return entity, db.HandleError(err)
}

func (r ExperimentRepository) Find(ctx context.Context, id int) (model.Experiment, error) {
	entity := model.Experiment{}
	query := fmt.Sprintf(
		"SELECT * FROM experiments WHERE id = $1 AND deleted_on IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, db.HandleError(err)
}

func (r ExperimentRepository) Create(ctx context.Context, entity *model.Experiment) error {
	query := `INSERT INTO experiments (name, description, status, created_on, updated_on)
				VALUES (:name, :description, :status, :created_on, :updated_on) RETURNING id;`
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

func (r ExperimentRepository) Update(ctx context.Context, entity model.Experiment) error {
	query := `UPDATE experiments
    		  	SET name = :name, 
    		  	    description = :description, 
    		  	    status = :status, 
    		  	    created_on = :created_on, 
    		  	    updated_on = :updated_on, 
    		  	    deleted_on = :deleted_on
				WHERE id = :id;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return db.HandleError(err)
}

func (r ExperimentRepository) FindAll(ctx context.Context) ([]model.Experiment, error) {
	var entities []model.Experiment
	query := fmt.Sprintf(
		"SELECT * FROM experiments WHERE deleted_on IS NULL",
	)
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, db.HandleError(err)
}
