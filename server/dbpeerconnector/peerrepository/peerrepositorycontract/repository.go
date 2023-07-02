package peerrepositorycontract

import (
	"context"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

type Repository interface {
	Save(ctx context.Context, name string, peer peerconnectorcontract.Peer) error
	Get(ctx context.Context, name string, execute func(peer peerconnectorcontract.Peer)) error
}
