package statisticsrepository

import (
	"cf-proposal/domain/datastore"
	"context"
)

var (
	MockGetStatisticEntity datastore.Statistic
	MockGetStatisticError  error
)

type StatisticsRepoMock struct{}

func (srm StatisticsRepoMock) GetStatistic(ctx context.Context, id int32) (datastore.Statistic, error) {
	return MockGetStatisticEntity, MockGetStatisticError
}
