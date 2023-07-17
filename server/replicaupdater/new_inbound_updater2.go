package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

type InboundUpdate struct {
	Key   string
	Value string
}

func NewFactory2(storage actormodel.ActorRef, recordUpdaterFactory RecordUpdaterFactory) replicaupdatercontract.Factory2 {
	return factory2{storage: storage, recordUpdaterFactory: recordUpdaterFactory}
}

type factory2 struct {
	storage              actormodel.ActorRef
	recordUpdaterFactory RecordUpdaterFactory
}

func (f factory2) NewInboundUpdater(ctx context.Context) (actormodel.ActorRef, error) {
	return actormodel.Spawn(
		ctx,
		&inboundUpdater{storage: f.storage, recordUpdaterFactory: f.recordUpdaterFactory},
		actormodel.WithBufferSize(16),
	), nil
}

type inboundUpdater struct {
	actormodel.Embed
	storage              actormodel.ActorRef
	recordUpdaterFactory RecordUpdaterFactory
}

func (u *inboundUpdater) Receive(sender actormodel.ActorRef, message any) actormodel.Actor {
	switch castedMsg := message.(type) {
	case InboundUpdate:
		return u.inboundUpdate(sender, castedMsg)
	}
	return u
}

func (u *inboundUpdater) inboundUpdate(_ actormodel.ActorRef, msg InboundUpdate) actormodel.Actor {
	recordUpdater := u.recordUpdaterFactory.New(nil, "", "")
	_ = u.TellBlocking(context.Background(), u.storage, dbstoragecontract.GetRecord{
		Key:     msg.Key,
		ReplyTo: recordUpdater,
	})
	return u
}
