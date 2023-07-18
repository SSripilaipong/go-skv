package recordreplicator

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func NewFactory(storage chan<- any, recordFactory dbstoragecontract.Factory) Factory {
	return factory{
		storage:       storage,
		recordFactory: recordFactory,
	}
}

type factory struct {
	storage       chan<- any
	recordFactory dbstoragecontract.Factory
}

func (f factory) New(ctx context.Context, key string, value string) (chan<- any, func()) {
	updater, join := actormodel.Spawn(ctx, &idle{
		storage:       f.storage,
		recordFactory: f.recordFactory,
		key:           key,
		value:         value,
	})
	updater <- commonmessage.Start{}
	return updater, join
}
