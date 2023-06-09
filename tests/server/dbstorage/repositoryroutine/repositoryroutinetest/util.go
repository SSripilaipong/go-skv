package repositoryroutinetest

import (
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/server/dbstorage/storagerecordfactory"
)

func NewStorageWithChannel(ch chan any) repositoryroutine.Interface {
	return repositoryroutine.New(ch, &RecordFactoryMock{})
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagerecordfactory.Interface) repositoryroutine.Interface {
	return repositoryroutine.New(ch, factory)
}
