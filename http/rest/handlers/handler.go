package handlers

import (
	"ab-testing/internal/experiment/repository"
	experimentService "ab-testing/internal/experiment/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	Logger            *logrus.Logger
	Router            *mux.Router
	ExperimentService experimentService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		Logger: lg,
		ExperimentService: experimentService.NewService(
			repository.NewExperimentRepository(db),
			repository.NewExperimentGroupRepository(db),
			repository.NewExperimentUserRepository(db),
		),
	}
}
