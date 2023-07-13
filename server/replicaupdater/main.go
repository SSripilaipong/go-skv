package replicaupdater

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewFactory(dbStorage dbstoragecontract.Storage) replicaupdatercontract.Factory {
	return factory{dbStorage: dbStorage}
}

type factory struct {
	dbStorage dbstoragecontract.Storage
}
