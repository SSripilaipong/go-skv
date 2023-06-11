package dbstorage

import (
	"go-skv/server/dbstorage/repositoryinteractor"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/server/dbstorage/storagerecord"
)

func New(storageBufferSize int, recordBufferSize int) (Repository, RepositoryInteractor) {
	ch := make(chan any, storageBufferSize)
	routine := repositoryroutine.New(ch, storagerecord.NewFactory(recordBufferSize))
	interactor := repositoryinteractor.New(ch)
	return routine, interactor
}
