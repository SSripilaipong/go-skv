package dbstorage

import (
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/server/dbstorage/storagerecordfactory"
)

func New(storageBufferSize int, recordBufferSize int) (Manager, chan<- any) {
	ch := make(chan any, storageBufferSize)
	s := storagemanager.New(ch, storagerecordfactory.NewFactory(recordBufferSize))
	return s, ch
}
