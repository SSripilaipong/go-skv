package storagerepository

import "go-skv/server/dbstorage/storagerecord"

type GetRecordSuccessCallback func(storagerecord.DbRecord)

type GetRecordMessage struct {
	Key     string
	Success GetRecordSuccessCallback
}