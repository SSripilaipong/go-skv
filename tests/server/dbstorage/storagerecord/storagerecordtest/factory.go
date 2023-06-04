package storagerecordtest

import (
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/server/dbstorage/storagerecord"
)

func NewFactory() storagemanager.RecordFactory {
	return storagerecord.NewFactory(1)
}
