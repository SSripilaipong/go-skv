package dbusecase

import (
	"go-skv/server/dbstorage/dbstoragecontract"
)

func New(repo dbstoragecontract.Storage) Interface {
	return usecase{repo: repo}
}

type usecase struct {
	repo dbstoragecontract.Storage
}
