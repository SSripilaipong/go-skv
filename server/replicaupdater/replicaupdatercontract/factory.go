package replicaupdatercontract

import "context"

type Factory interface {
	NewInboundUpdater(ctx context.Context) (InboundUpdater, error)
}

type InboundUpdater interface {
	Update(key string, value string) error
	Join()
}
