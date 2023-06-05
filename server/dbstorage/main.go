package dbstorage

import (
	"go-skv/server/dbstorage/storagerecordfactory"
	"go-skv/server/dbstorage/storagerepository/repositoryroutine"
)

func New(storageBufferSize int, recordBufferSize int) (Repository, chan<- any) {
	ch := make(chan any, storageBufferSize)
	s := repositoryroutine.New(ch, storagerecordfactory.NewFactory(recordBufferSize))
	return s, ch
}
