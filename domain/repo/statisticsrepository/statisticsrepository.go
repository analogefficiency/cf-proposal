package statisticsrepository

import (
	"cf-proposal/common/helper"
	"cf-proposal/domain/datastore"
	"context"
	"database/sql"
)

type StatisticsInterface interface {
	GetStatistic(ctx context.Context, id int32) (datastore.Statistic, error)
}

type Statistic struct {
	repo StatisticsInterface
}

func Init(repo StatisticsInterface) *Statistic {
	return &Statistic{
		repo: repo,
	}
}

func (s *Statistic) GetStatistic(ctx context.Context, id int) (datastore.Statistic, error) {

	statistic, err := s.repo.GetStatistic(ctx, int32(id))
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return datastore.Statistic{
				UrlID:           int32(id),
				TwentyFourHours: 0,
				LastSevenDays:   0,
				AllTime:         0,
			}, nil
		}
		return datastore.Statistic{}, &helper.CustomError{Message: err.Error()}
	}
	return statistic, nil
}
