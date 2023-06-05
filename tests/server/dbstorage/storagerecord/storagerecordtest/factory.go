package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecordfactory"
)

func NewFactory() storagerecordfactory.Interface {
	return storagerecordfactory.NewFactory(1)
}
