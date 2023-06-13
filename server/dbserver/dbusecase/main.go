package dbusecase

import "go-skv/server/dbstorage"

func New(repo dbstorage.RepositoryInteractor) Interface {
	return usecase{repo: repo}
}

type usecase struct {
	repo dbstorage.RepositoryInteractor
}
