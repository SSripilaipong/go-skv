package connectormanager

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

func New(existingPeerAddresses []string, client peerclientcontract.Client, peerRepo peerrepositorycontract.Repository) peerconnectorcontract.Connector {
	return manager{
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
		peerRepo:              peerRepo,
	}
}

type manager struct {
	existingPeerAddresses []string
	client                peerclientcontract.Client
	peerRepo              peerrepositorycontract.Repository
}

var _ peerconnectorcontract.Connector = manager{}
