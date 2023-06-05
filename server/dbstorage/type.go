package dbstorage

import (
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/server/dbstorage/storagerecord"
)

type Manager = storagemanager.Interface
type Record = storagerecord.DbRecord
type GetValueResponse = storagerecord.GetValueResponse
type SetValueResponse = storagerecord.SetValueResponse
