package replicaupdatercontract

import (
	"context"
	"go-skv/common/actormodel"
)

type Factory interface {
	NewInboundUpdater(ctx context.Context) (InboundUpdater, error)
}

type InboundUpdater interface {
	Update(key string, value string) error
	Join()
}

type Factory2 interface {
	NewInboundUpdater(ctx context.Context) (actormodel.ActorRef, error)
}
