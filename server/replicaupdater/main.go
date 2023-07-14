package replicaupdater

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewFactory(dbStorage dbstoragecontract.Storage, recordService RecordService, recordFactory dbstoragecontract.Factory) replicaupdatercontract.Factory {
	return factory{
		dbStorage:     dbStorage,
		recordService: recordService,
		recordFactory: recordFactory,
	}
}

type factory struct {
	dbStorage     dbstoragecontract.Storage
	recordService RecordService
	recordFactory dbstoragecontract.Factory
}
