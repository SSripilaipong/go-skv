package storagerepositorytest

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/tests/server/dbstorage/dbstoragetest"
)

type RecordFactoryMock struct {
	New_Return   dbstoragecontract.Record
	New_IsCalled bool
	New_ctx      context.Context
}

func (f *RecordFactoryMock) New(ctx context.Context) dbstoragecontract.Record {
	f.New_IsCalled = true
	f.New_ctx = ctx
	return goutil.Coalesce[dbstoragecontract.Record](f.New_Return, &dbstoragetest.RecordMock{})
}

func (f *RecordFactoryMock) New_CaptureReset() {
	f.New_IsCalled = false
}
