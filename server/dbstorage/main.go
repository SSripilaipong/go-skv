package dbstorage

import (
	"go-skv/server/dbstorage/repositoryinteractor"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/server/dbstorage/storagerecord"
)

func New(storageBufferSize int, recordBufferSize int) (Repository, chan<- any) { // TODO: return interactor instead of channel
	s, _, ch := newRepository(storageBufferSize, storagerecord.NewFactory(recordBufferSize))
	return s, ch
}

func newRepository(storageBufferSize int, factory storagerecord.Factory) (repositoryroutine.Interface, repositoryinteractor.Interface, chan<- any) {
	ch := make(chan any, storageBufferSize)
	routine := repositoryroutine.New(ch, factory)
	interactor := repositoryinteractor.New(ch)
	return routine, interactor, ch
}
