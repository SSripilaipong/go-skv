package connectormanager

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

func New(ctx context.Context, existingPeerAddresses []string, client peerclientcontract.Client, peerRepo peerrepositorycontract.Repository) peerconnectorcontract.Connector {
	subCtx, cancelSubCtx := context.WithCancel(ctx)
	return manager{
		ctx:                   ctx,
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
		peerRepo:              peerRepo,

		subCtx:       subCtx,
		cancelSubCtx: cancelSubCtx,
	}
}

type manager struct {
	ctx                   context.Context
	existingPeerAddresses []string
	client                peerclientcontract.Client
	peerRepo              peerrepositorycontract.Repository

	subCtx       context.Context
	cancelSubCtx context.CancelFunc
}

var _ peerconnectorcontract.Connector = manager{}
