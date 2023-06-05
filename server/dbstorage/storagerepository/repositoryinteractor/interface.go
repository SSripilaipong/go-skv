package repositoryinteractor

import (
	"go-skv/server/dbstorage/storagerepository/repositoryroutine"
	"time"
)

type Interface interface {
	GetRecord(key string, success repositoryroutine.GetRecordSuccessCallback, timeout time.Duration) error
	GetOrCreateRecord(key string, success repositoryroutine.GetOrCreateRecordSuccessCallback, timeout time.Duration) error
}
