package dbstorage

import (
	"go-skv/server/dbstorage/storagerecordfactory"
	"go-skv/server/dbstorage/storagerepository"
)

func New(storageBufferSize int, recordBufferSize int) (Repository, chan<- any) { // TODO: return interactor instead of channel
	s, _, ch := storagerepository.New(storageBufferSize, storagerecordfactory.NewFactory(recordBufferSize))
	return s, ch
}
