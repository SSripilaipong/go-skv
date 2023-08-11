package replicaupdater

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/recordreplicator"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewFactory2(dbStorage dbstoragecontract.Storage, recordFactory dbstoragecontract.Factory) replicaupdatercontract.Factory {
	return NewFactoryAdapter(NewActorFactory(recordreplicator.NewFactory(newStorageAdapter(dbStorage), recordFactory)))
}
