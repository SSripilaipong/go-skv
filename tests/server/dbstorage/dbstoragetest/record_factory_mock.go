package dbstoragetest

import (
	"go-skv/server/dbstorage"
	"go-skv/util/goutil"
)

type RecordFactoryMock struct {
	New_Return   dbstorage.DbRecord
	New_IsCalled bool
}

func (f *RecordFactoryMock) New() dbstorage.DbRecord {
	f.New_IsCalled = true
	return goutil.Coalesce[dbstorage.DbRecord](f.New_Return, &RecordMock{})
}

func (f *RecordFactoryMock) New_CaptureReset() {
	f.New_IsCalled = false
}
