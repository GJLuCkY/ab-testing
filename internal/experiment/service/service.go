package service

import (
	"ab-testing/internal/experiment/repository"
)

type Service struct {
	experimentRepository      repository.ExperimentRepository
	experimentGroupRepository repository.ExperimentGroupRepository
	experimentUserRepository  repository.ExperimentUserRepository
}

func NewService(
	experimentRepository repository.ExperimentRepository,
	experimentGroupRepository repository.ExperimentGroupRepository,
	experimentUserRepository repository.ExperimentUserRepository,
) Service {
	return Service{
		experimentRepository:      experimentRepository,
		experimentGroupRepository: experimentGroupRepository,
		experimentUserRepository:  experimentUserRepository,
	}
}
