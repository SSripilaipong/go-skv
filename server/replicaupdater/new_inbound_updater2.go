package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/server/replicaupdater/recordreplicator"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

type InboundUpdate struct {
	Key   string
	Value string
}

func NewFactory2(recordUpdaterFactory recordreplicator.Factory) replicaupdatercontract.Factory2 {
	return factory2{recordUpdaterFactory: recordUpdaterFactory}
}

type factory2 struct {
	recordUpdaterFactory recordreplicator.Factory
}

func (f factory2) NewInboundUpdater(ctx context.Context) (chan<- any, error) {
	ch, _ := actormodel.Spawn(
		ctx,
		&inboundUpdater{recordUpdaterFactory: f.recordUpdaterFactory},
		actormodel.WithBufferSize(16),
	)
	return ch, nil
}

type inboundUpdater struct {
	actormodel.Embed
	recordUpdaterFactory recordreplicator.Factory
}

func (u *inboundUpdater) Receive(message any) actormodel.Actor {
	switch castedMsg := message.(type) {
	case InboundUpdate:
		return u.inboundUpdate(castedMsg)
	}
	return u
}

func (u *inboundUpdater) inboundUpdate(msg InboundUpdate) actormodel.Actor {
	_, _ = u.recordUpdaterFactory.New(u.Ctx(), msg.Key, msg.Value)
	return u
}