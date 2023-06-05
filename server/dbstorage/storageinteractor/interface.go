package storageinteractor

import (
	"go-skv/server/dbstorage/storagerepository"
	"time"
)

type Interface interface {
	GetRecord(key string, success storagerepository.GetRecordSuccessCallback, timeout time.Duration) error
	GetOrCreateRecord(key string, success storagerepository.GetOrCreateRecordSuccessCallback, timeout time.Duration) error
}
