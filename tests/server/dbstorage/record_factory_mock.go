package dbstorageTest

import (
	"go-skv/goutil"
	"go-skv/server/dbstorage"
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
