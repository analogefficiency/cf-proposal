package repository

import (
	"cf-proposal/domain/model"
	"context"
	"database/sql"
	"fmt"
)

type StatisticsRepo struct {
	q *Queries
}

func InitStatisticsRepo(db *sql.DB) *StatisticsRepo {
	return &StatisticsRepo{
		q: New(db),
	}
}

func (s *StatisticsRepo) GetStatistic(ctx context.Context, id int32) (model.StatisticsDto, error) {
	stats, err := s.q.GetStatisticsByUrl(ctx, id)
	if err != nil {
		return model.StatisticsDto{}, fmt.Errorf("%w", err)
	}
	return model.StatisticsDto{
		UrlID:           stats.UrlID,
		TwentyFourHours: stats.TwentyFourHours,
		LastSevenDays:   stats.LastSevenDays,
		AllTime:         stats.AllTime,
	}, nil
}
