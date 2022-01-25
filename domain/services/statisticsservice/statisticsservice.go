package statisticsservice

import (
	"cf-proposal/domain/model"
	"context"
	"fmt"
	"strconv"
)

type StatisticsRepo interface {
	GetStatistic(ctx context.Context, id int32) (model.StatisticsDto, error)
}

type Statistic struct {
	repo StatisticsRepo
}

func Init(repo StatisticsRepo) *Statistic {
	return &Statistic{
		repo: repo,
	}
}

func (s *Statistic) GetStatistic(ctx context.Context, id string) (model.StatisticsDto, error) {

	convertedId, convErr := strconv.Atoi(id)
	if convErr != nil {
		return model.StatisticsDto{}, fmt.Errorf("%w", convErr)
	}
	longUrl, repoErr := s.repo.GetStatistic(ctx, int32(convertedId))
	if repoErr != nil {
		return model.StatisticsDto{}, fmt.Errorf("%w", repoErr)
	}
	return longUrl, nil
}
