package repositoryroutinetest

import (
	"go-skv/server/dbstorage/storagerecordfactory"
	"go-skv/server/dbstorage/storagerepository/repositoryroutine"
)

func NewStorageWithChannel(ch chan any) repositoryroutine.Interface {
	return repositoryroutine.New(ch, nil)
}

func NewStorageWithChannelAndRecordFactory(ch chan any, factory storagerecordfactory.Interface) repositoryroutine.Interface {
	return repositoryroutine.New(ch, factory)
}
