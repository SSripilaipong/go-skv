package dbstorage

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/dbstorage/storagerepository"
)

func New(storageBufferSize int, recordBufferSize int) Repository {
	recordFactory := storagerecord.NewFactory(recordBufferSize)
	return storagerepository.New(storageBufferSize, recordFactory)
}
