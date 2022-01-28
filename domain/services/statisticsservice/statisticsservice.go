package statisticsservice

import (
	"cf-proposal/common/helper"
	"cf-proposal/common/messages"
	"cf-proposal/domain/model"
	"context"
	"database/sql"
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

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return model.StatisticsDto{}, &helper.CustomError{Message: fmt.Sprintf(messages.TYPE_MISMATCH, "id", "string", "int")}
	}
	statistic, err := s.repo.GetStatistic(ctx, int32(convertedId))
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return model.StatisticsDto{
				UrlID:           int32(convertedId),
				TwentyFourHours: 0,
				LastSevenDays:   0,
				AllTime:         0,
			}, nil
		}
		return model.StatisticsDto{}, &helper.CustomError{Message: err.Error()}
	}
	return statistic, nil
}
