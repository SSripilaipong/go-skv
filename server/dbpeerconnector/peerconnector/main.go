package peerconnector

import (
	"context"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

func New(ctx context.Context, existingPeerAddresses []string, client peerclientcontract.Client, peerRepo peerrepositorycontract.Repository) peerconnectorcontract.Connector {
	return connector{
		ctx:                   ctx,
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
		peerRepo:              peerRepo,
	}
}

type connector struct {
	ctx                   context.Context
	existingPeerAddresses []string
	client                peerclientcontract.Client
	peerRepo              peerrepositorycontract.Repository
}

var _ peerconnectorcontract.Connector = connector{}
