package peerclientcontract

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type Client interface {
	ConnectToPeer(ctx context.Context, address string) (peerconnectorcontract.Peer, error)
	WaitForAllToBeDisconnected() error
}
