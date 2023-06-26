package connectormanager

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

func New(ctx context.Context, existingPeerAddresses []string, client peerclientcontract.Client, peerRepo peerrepositorycontract.Repository) peerconnectorcontract.Connector {
	return manager{
		ctx:                   ctx,
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
		peerRepo:              peerRepo,
	}
}

type manager struct {
	ctx                   context.Context
	existingPeerAddresses []string
	client                peerclientcontract.Client
	peerRepo              peerrepositorycontract.Repository
}

var _ peerconnectorcontract.Connector = manager{}
