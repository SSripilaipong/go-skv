package replicaupdater

import (
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/recordreplicator"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewFactory(dbStorage dbstoragecontract.Storage, recordFactory dbstoragecontract.Factory) replicaupdatercontract.ActorFactory {
	return NewActorFactory(recordreplicator.NewFactory(newStorageAdapter(dbStorage), recordFactory))
}
