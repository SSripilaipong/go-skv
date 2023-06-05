package storageinteractor

import (
	"go-skv/server/dbstorage/storagerepository"
	"time"
)

type Interface interface {
	GetRecord(key string, callback storagerepository.GetRecordSuccessCallback, timeout time.Duration) error
}
