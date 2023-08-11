package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/server/replicaupdater/recordreplicator"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewActorFactory(recordUpdaterFactory recordreplicator.Factory) replicaupdatercontract.ActorFactory {
	return factory{recordUpdaterFactory: recordUpdaterFactory}
}

type factory struct {
	recordUpdaterFactory recordreplicator.Factory
}

func (f factory) NewInboundUpdater(ctx context.Context) (chan<- any, error) {
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
	case replicaupdatercontract.InboundUpdate:
		return u.inboundUpdate(castedMsg)
	}
	return u
}

func (u *inboundUpdater) inboundUpdate(msg replicaupdatercontract.InboundUpdate) actormodel.Actor {
	_, _ = u.recordUpdaterFactory.New(u.Ctx(), msg.Key, msg.Value)
	return u
}
