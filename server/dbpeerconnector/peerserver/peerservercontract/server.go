package peerservercontract

import "context"

type Server interface {
	Start(ctx context.Context) error
	Port() int
	Join()
}
