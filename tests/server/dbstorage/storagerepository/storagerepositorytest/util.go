package storagerepositorytest

import (
	"go-skv/server/dbstorage/storagerecordfactory"
	"go-skv/server/dbstorage/storagerepository"
)

func NewStorageWithChannel(ch chan any) storagerepository.Interface {
	return storagerepository.New(ch, nil)
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagerecordfactory.Interface) storagerepository.Interface {
	return storagerepository.New(ch, factory)
}
