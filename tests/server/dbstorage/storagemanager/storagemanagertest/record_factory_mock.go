package storagemanagertest

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/util/goutil"
)

type RecordFactoryMock struct {
	New_Return   storagerecord.DbRecord
	New_IsCalled bool
	New_ctx      context.Context
}

func (f *RecordFactoryMock) New(ctx context.Context) storagerecord.DbRecord {
	f.New_IsCalled = true
	f.New_ctx = ctx
	return goutil.Coalesce[storagerecord.DbRecord](f.New_Return, &RecordMock{})
}

func (f *RecordFactoryMock) New_CaptureReset() {
	f.New_IsCalled = false
}
