package clientsidepeercontract

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type Factory interface {
	New(ctx context.Context) (peerconnectorcontract.Peer, error)
}
