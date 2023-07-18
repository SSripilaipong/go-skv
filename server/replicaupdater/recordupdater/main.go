package recordupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
)

func NewFactory(storage chan<- any) Factory {
	return factory{storage: storage}
}

type factory struct {
	storage chan<- any
}

func (f factory) New(ctx context.Context, key string, value string) chan<- any {
	updater, _ := actormodel.Spawn(ctx, &idleState{
		storage: f.storage,
		key:     key,
		value:   value,
	})
	updater <- commonmessage.Start{}
	return updater
}
