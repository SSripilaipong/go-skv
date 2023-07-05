package dbstorage

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
)

func New(storageBufferSize int, recordBufferSize int) dbstoragecontract.Storage {
	recordFactory := storagerecord.NewFactory(recordBufferSize)
	return storagerepository.New(storageBufferSize, recordFactory)
}
