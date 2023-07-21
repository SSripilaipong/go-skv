package replicaupdater

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/recordreplicator"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewFactory(dbStorage dbstoragecontract.Storage, recordService RecordService, recordFactory dbstoragecontract.Factory) replicaupdatercontract.Factory {
	return factory{
		dbStorage:     dbStorage,
		recordService: recordService,
		recordFactory: recordFactory,
	}
}

func NewFactory2(dbStorage dbstoragecontract.Storage, recordFactory dbstoragecontract.Factory) replicaupdatercontract.Factory {
	return NewFactoryAdapter(NewActorFactory(recordreplicator.NewFactory(newStorageAdapter(dbStorage), recordFactory)))
}

type factory struct {
	dbStorage     dbstoragecontract.Storage
	recordService RecordService
	recordFactory dbstoragecontract.Factory
}
