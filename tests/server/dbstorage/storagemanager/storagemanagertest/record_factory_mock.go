package storagemanagertest

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/util/goutil"
)

type RecordFactoryMock struct {
	New_Return   storagemanager.DbRecord
	New_IsCalled bool
}

func (f *RecordFactoryMock) New(context.Context) storagemanager.DbRecord {
	f.New_IsCalled = true
	return goutil.Coalesce[storagemanager.DbRecord](f.New_Return, &RecordMock{})
}

func (f *RecordFactoryMock) New_CaptureReset() {
	f.New_IsCalled = false
}
