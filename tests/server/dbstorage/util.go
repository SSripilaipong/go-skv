package dbstorageTest

import "go-skv/server/dbstorage"

func NewStorageWithChannel(ch chan any) dbstorage.Interface {
	return dbstorage.New(ch, nil)
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory dbstorage.RecordFactory) dbstorage.Interface {
	return dbstorage.New(ch, factory)
}
