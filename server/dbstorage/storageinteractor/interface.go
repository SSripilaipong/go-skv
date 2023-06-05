package storageinteractor

import (
	"go-skv/server/dbstorage/storagerepository"
)

type Interface interface {
	GetRecord(key string, callback storagerepository.GetRecordSuccessCallback) error
}
