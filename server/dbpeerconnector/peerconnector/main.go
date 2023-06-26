package peerconnector

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
)

func New(existingPeerAddresses []string, client peerclientcontract.Client, peerRepo peerrepositorycontract.Repository) peerconnectorcontract.Connector {
	return connector{
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
		peerRepo:              peerRepo,
	}
}

type connector struct {
	existingPeerAddresses []string
	client                peerclientcontract.Client
	peerRepo              peerrepositorycontract.Repository
}

var _ peerconnectorcontract.Connector = connector{}
