package dbusecasetest

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/repositoryroutine"
)

type RepoMock struct {
	GetRecord_key string
	GetRecord_ctx context.Context
}

var _ dbstorage.RepositoryInteractor = &RepoMock{}

func (r *RepoMock) GetRecord(ctx context.Context, key string, success repositoryroutine.GetRecordSuccessCallback) error {
	r.GetRecord_key = key
	r.GetRecord_ctx = ctx
	return nil
}

func (r *RepoMock) GetOrCreateRecord(ctx context.Context, key string, success repositoryroutine.GetOrCreateRecordSuccessCallback) error {
	//TODO implement me
	panic("implement me")
}
