package dbstorage

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
)

type Repository = storagerepository.Interface
type Record = storagerecord.DbRecord
type GetValueResponse = storagerecord.GetValueResponse
type SetValueResponse = storagerecord.SetValueResponse
