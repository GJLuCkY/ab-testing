package service

import (
	"context"
	"errors"
	"sort"
	"strconv"

	"ab-testing/internal/experiment/model"
	"ab-testing/pkg/db"
	"ab-testing/pkg/erru"
)

func (s Service) GetBySlug(ctx context.Context, slug string) (model.Experiment, error) {
	experiment, err := s.experimentRepository.FirstBySlug(ctx, slug, true)

	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.Experiment{}, erru.ErrArgument{errors.New("Эксперимент не найден, либо он выключен")}
	default:
		return model.Experiment{}, err
	}

	return experiment, nil
}

func (s Service) GetActiveGroups(ctx context.Context, experimentId int) ([]model.ExperimentGroup, error) {
	experimentGroups, err := s.experimentGroupRepository.GetByExperimentId(ctx, experimentId, true)

	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []model.ExperimentGroup{}, erru.ErrArgument{errors.New("Эксперимент не найден, либо он выключен")}
	default:
		return []model.ExperimentGroup{}, err
	}

	return experimentGroups, nil
}

func (s Service) GetGroup(ctx context.Context, experimentId int, userID string) (model.ExperimentGroup, error) {
	experimentUser, err := s.experimentUserRepository.GetByUserIdAndExperimentId(ctx, experimentId, userID)

	var groupId int

	if errors.As(err, &db.ErrObjectNotFound{}) {
		groups, err := s.experimentGroupRepository.GetActiveGroupsWithUserCount(ctx, experimentId)
		if err != nil {
			return model.ExperimentGroup{}, err
		}

		if len(groups) > 1 {
			groupId = s.AssignUserToGroup(groups)
		} else {
			groupId = groups[0].ID
		}

		userIDInt, _ := strconv.Atoi(userID)

		experimentUser := model.ExperimentUser{
			ExperimentID: experimentId,
			GroupID:      groupId,
			UserID:       userIDInt,
		}

		err = s.experimentUserRepository.Create(ctx, experimentUser)
		if err != nil {
			return model.ExperimentGroup{}, err
		}
	} else {
		groupId = experimentUser.GroupID
	}

	experimentGroup, err := s.experimentGroupRepository.Find(ctx, groupId)

	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.ExperimentGroup{}, erru.ErrArgument{errors.New("При распределении пользователя в группу произошла ошибка")}
	default:
		return model.ExperimentGroup{}, err
	}

	return experimentGroup, nil
}

func (s Service) AssignUserToGroup(groups []model.ExperimentGroup) int {
	for i := range groups {
		group := &groups[i]
		group.OccupancyRatio = group.CoveragePercentage / float64(group.UserCount)
	}

	compare := func(i, j int) bool {
		return groups[i].OccupancyRatio > groups[j].OccupancyRatio
	}

	sort.Slice(groups, compare)

	return groups[0].ID
}
