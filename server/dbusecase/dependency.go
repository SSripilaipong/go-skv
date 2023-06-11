package dbusecase

import (
	"go-skv/server/dbstorage"
)

type Dependency struct {
	repo dbstorage.RepositoryInteractor
}

func NewDependency(repo dbstorage.RepositoryInteractor) Dependency {
	return Dependency{
		repo: repo,
	}
}
