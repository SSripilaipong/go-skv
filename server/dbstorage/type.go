package dbstorage

import (
	"go-skv/server/dbstorage/repositoryinteractor"
	"go-skv/server/dbstorage/repositoryroutine"
	"go-skv/server/dbstorage/storagerecord"
)

type Repository = repositoryroutine.Interface
type RepositoryInteractor = repositoryinteractor.Interface
type Record = storagerecord.Interface
type GetValueResponse = storagerecord.GetValueResponse
type SetValueResponse = storagerecord.SetValueResponse
