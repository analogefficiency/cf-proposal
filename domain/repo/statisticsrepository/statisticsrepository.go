package statisticsrepository

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"cf-proposal/domain/datastore"
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func (s *Statistic) GetStatistic(ctx context.Context, id string) (datastore.Statistic, error) {

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return datastore.Statistic{}, &helper.CustomError{Message: fmt.Sprintf(messages.TYPE_MISMATCH, "id", "string", "int")}
	}
	statistic, err := s.repo.GetStatistic(ctx, int32(convertedId))
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return datastore.Statistic{
				UrlID:           int32(convertedId),
				TwentyFourHours: 0,
				LastSevenDays:   0,
				AllTime:         0,
			}, nil
		}
		return datastore.Statistic{}, &helper.CustomError{Message: err.Error()}
	}
	return statistic, nil
}
