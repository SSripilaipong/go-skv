package storagerepository

import (
	"go-skv/server/dbstorage/storagerecordfactory"
	"go-skv/server/dbstorage/storagerepository/repositoryinteractor"
	"go-skv/server/dbstorage/storagerepository/repositoryroutine"
)

func New(storageBufferSize int, factory storagerecordfactory.Interface) (Routine, Interactor, chan<- any) {
	ch := make(chan any, storageBufferSize)
	routine := repositoryroutine.New(ch, factory)
	interactor := repositoryinteractor.New(ch)
	return routine, interactor, ch
}
