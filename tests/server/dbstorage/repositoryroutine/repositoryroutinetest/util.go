package repositoryroutinetest

import (
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/server/dbstorage/storagerecord"
)

func NewStorageWithChannel(ch chan any) repositoryroutine.Interface {
	return repositoryroutine.New(ch, &RecordFactoryMock{})
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagerecord.Factory) repositoryroutine.Interface {
	return repositoryroutine.New(ch, factory)
}
