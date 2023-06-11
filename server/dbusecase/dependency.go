package dbusecase

import (
	"go-skv/server/dbstorage"
)

type Dependency struct {
	storageChan chan<- any
	repo        dbstorage.RepositoryInteractor
}

func NewDependency(storageChan chan<- any, repo dbstorage.RepositoryInteractor) Dependency {
	return Dependency{
		storageChan: storageChan,
		repo:        repo,
	}
}
