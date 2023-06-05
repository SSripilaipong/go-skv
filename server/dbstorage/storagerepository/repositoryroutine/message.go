package repositoryroutine

import "go-skv/server/dbstorage/storagerecord"

type GetRecordSuccessCallback func(storagerecord.DbRecord)

type GetOrCreateRecordSuccessCallback func(storagerecord.DbRecord)

type GetRecordMessage struct {
	Key     string
	Success GetRecordSuccessCallback
}

type GetOrCreateRecordMessage struct {
	Key     string
	Success GetOrCreateRecordSuccessCallback
}
