package repositoryroutine

import "go-skv/server/dbstorage/storagerecord"

type GetRecordSuccessCallback func(storagerecord.Interface)

type GetOrCreateRecordSuccessCallback func(storagerecord.Interface)

type GetRecordMessage struct {
	Key     string
	Success GetRecordSuccessCallback
}

type GetOrCreateRecordMessage struct {
	Key     string
	Success GetOrCreateRecordSuccessCallback
}
