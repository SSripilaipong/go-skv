package dbusecasetest

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/repositoryroutine"
)

type RepoMock struct {
	GetRecord_key string
}

var _ dbstorage.RepositoryInteractor = &RepoMock{}

func (r *RepoMock) GetRecord(ctx context.Context, key string, success repositoryroutine.GetRecordSuccessCallback) error {
	r.GetRecord_key = key
	return nil
}

func (r *RepoMock) GetOrCreateRecord(ctx context.Context, key string, success repositoryroutine.GetOrCreateRecordSuccessCallback) error {
	//TODO implement me
	panic("implement me")
}
