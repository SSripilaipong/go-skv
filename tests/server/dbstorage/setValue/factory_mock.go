package setValue

import "go-skv/server/dbstorage"

type recordFactoryMock struct {
	New_Return   dbstorage.DbRecord
	New_IsCalled bool
}

func (f *recordFactoryMock) New() dbstorage.DbRecord {
	f.New_IsCalled = true
	return f.New_Return
}
