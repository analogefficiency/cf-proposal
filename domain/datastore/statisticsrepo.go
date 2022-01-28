package datastore

import (
	"context"
	"database/sql"
	"fmt"
)

type StatisticsRepo struct {
	q *Queries
}

func InitStatisticsDatastore(db *sql.DB) *StatisticsRepo {
	return &StatisticsRepo{
		q: New(db),
	}
}

func (s *StatisticsRepo) GetStatistic(ctx context.Context, id int32) (Statistic, error) {
	stats, err := s.q.GetStatisticsByUrl(ctx, id)
	if err != nil {
		return Statistic{}, fmt.Errorf("%w", err)
	}
	return Statistic{
		UrlID:           stats.UrlID,
		TwentyFourHours: stats.TwentyFourHours,
		LastSevenDays:   stats.LastSevenDays,
		AllTime:         stats.AllTime,
	}, nil
}
