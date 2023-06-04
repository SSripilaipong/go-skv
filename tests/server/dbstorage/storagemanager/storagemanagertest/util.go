package storagemanagertest

import (
	"go-skv/server/dbstorage/storagemanager"
)

func NewStorageWithChannel(ch chan any) storagemanager.Interface {
	return storagemanager.New(ch, nil)
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagemanager.RecordFactory) storagemanager.Interface {
	return storagemanager.New(ch, factory)
}
