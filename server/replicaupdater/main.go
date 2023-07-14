package replicaupdater

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewFactory(dbStorage dbstoragecontract.Storage, recordService RecordService) replicaupdatercontract.Factory {
	return factory{
		dbStorage:     dbStorage,
		recordService: recordService,
	}
}

type factory struct {
	dbStorage     dbstoragecontract.Storage
	recordService RecordService
}
