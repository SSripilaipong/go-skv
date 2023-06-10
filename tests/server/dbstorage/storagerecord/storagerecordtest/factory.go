package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecord"
)

func NewFactory() storagerecord.Factory {
	return storagerecord.NewFactory(1)
}
