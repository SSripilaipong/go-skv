package storagerepositorytest

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/dbstoragetest"
)

type RecordFactoryMock struct {
	New_Return   storagerecord.Interface
	New_IsCalled bool
	New_ctx      context.Context
}

func (f *RecordFactoryMock) New(ctx context.Context) storagerecord.Interface {
	f.New_IsCalled = true
	f.New_ctx = ctx
	return goutil.Coalesce[storagerecord.Interface](f.New_Return, &dbstoragetest.RecordMock{})
}

func (f *RecordFactoryMock) New_CaptureReset() {
	f.New_IsCalled = false
}
