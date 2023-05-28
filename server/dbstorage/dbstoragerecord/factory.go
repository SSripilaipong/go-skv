package dbstoragerecord

import "go-skv/server/dbstorage"

type recordFactory struct{}

func NewFactory() dbstorage.RecordFactory {
	return &recordFactory{}
}

func (r *recordFactory) New() dbstorage.DbRecord {
	return newRecordInterface()
}
