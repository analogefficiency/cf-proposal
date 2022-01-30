package historyrepository

import (
	"cf-proposal/domain/datastore"
	"context"
)

var (
	MockCreateHistoryDto datastore.History
	MockDeletedRows      int64
	MockCreateError      error
	MockDeleteError      error
)

type HistoryRepoMock struct{}

func (hrm HistoryRepoMock) Create(ctx context.Context, id int32) (datastore.History, error) {
	return MockCreateHistoryDto, MockCreateError
}
func (hrm HistoryRepoMock) Delete(ctx context.Context, id int32) (int64, error) {
	return MockDeletedRows, MockDeleteError
}
