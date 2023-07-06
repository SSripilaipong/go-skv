package replicaupdatercontract

type Factory interface {
	NewInboundUpdater() (InboundUpdater, error)
}

type InboundUpdater interface {
}
