package peerconnectorcontract

import "context"

type Connector interface {
	Start(ctx context.Context) error
	Join() error
}
