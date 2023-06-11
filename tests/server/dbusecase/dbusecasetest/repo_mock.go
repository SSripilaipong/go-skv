package dbusecasetest

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/tests/server/dbstorage/dbstoragetest"
	"go-skv/util/goutil"
)

type RepoMock struct {
	GetRecord_key            string
	GetRecord_ctx            context.Context
	GetRecord_success_record dbstorage.Record
	GetOrCreateRecord_key    string
}

var _ dbstorage.RepositoryInteractor = &RepoMock{}

func (r *RepoMock) GetRecord(ctx context.Context, key string, success repositoryroutine.GetRecordSuccessCallback) error {
	r.GetRecord_key = key
	r.GetRecord_ctx = ctx
	go success(goutil.Coalesce[dbstorage.Record](r.GetRecord_success_record, &dbstoragetest.RecordMock{}))
	return nil
}

func (r *RepoMock) GetOrCreateRecord(ctx context.Context, key string, success repositoryroutine.GetOrCreateRecordSuccessCallback) error {
	r.GetOrCreateRecord_key = key
	return nil
}
