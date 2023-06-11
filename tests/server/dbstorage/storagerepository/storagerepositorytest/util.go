package storagerepositorytest

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
)

func NewStorageWithChannel(ch chan any) storagerepository.Interface {
	return storagerepository.NewRepository(ch, &RecordFactoryMock{})
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagerecord.Factory) storagerepository.Interface {
	return storagerepository.NewRepository(ch, factory)
}
