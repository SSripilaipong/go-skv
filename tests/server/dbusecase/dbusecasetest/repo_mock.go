package dbusecasetest

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/storagerepository"
	"go-skv/tests/server/dbstorage/dbstoragetest"
)

type RepoMock struct {
	GetRecord_key                    string
	GetRecord_ctx                    context.Context
	GetRecord_success_record         dbstorage.Record
	GetOrCreateRecord_key            string
	GetOrCreateRecord_ctx            context.Context
	GetOrCreateRecord_success_record dbstorage.Record
}

var _ dbstorage.RepositoryInteractor = &RepoMock{}

func (r *RepoMock) GetRecord(ctx context.Context, key string, success storagerepository.GetRecordSuccessCallback) error {
	r.GetRecord_key = key
	r.GetRecord_ctx = ctx
	go success(goutil.Coalesce[dbstorage.Record](r.GetRecord_success_record, &dbstoragetest.RecordMock{}))
	return nil
}

func (r *RepoMock) GetOrCreateRecord(ctx context.Context, key string, success storagerepository.GetOrCreateRecordSuccessCallback) error {
	r.GetOrCreateRecord_ctx = ctx
	r.GetOrCreateRecord_key = key
	go success(goutil.Coalesce[dbstorage.Record](r.GetOrCreateRecord_success_record, &dbstoragetest.RecordMock{}))
	return nil
}
