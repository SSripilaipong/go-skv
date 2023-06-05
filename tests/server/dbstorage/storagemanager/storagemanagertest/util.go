package storagemanagertest

import (
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/server/dbstorage/storagerecordfactory"
)

func NewStorageWithChannel(ch chan any) storagemanager.Interface {
	return storagemanager.New(ch, nil)
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagerecordfactory.Interface) storagemanager.Interface {
	return storagemanager.New(ch, factory)
}
