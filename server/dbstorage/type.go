package dbstorage

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository/repositoryroutine"
)

type Repository = repositoryroutine.Interface
type Record = storagerecord.DbRecord
type GetValueResponse = storagerecord.GetValueResponse
type SetValueResponse = storagerecord.SetValueResponse
