package dbstorage

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
)

type Repository = storagerepository.Interface
type RepositoryInteractor = storagerepository.Interactor
type Record = storagerecord.Interface
type GetValueResponse = storagerecord.GetValueResponse
type SetValueResponse = storagerecord.SetValueResponse
