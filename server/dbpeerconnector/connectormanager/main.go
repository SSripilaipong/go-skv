package connectormanager

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
	"go-skv/server/dbpeerconnector/peerserver/peerservercontract"
)

func New(existingPeerAddresses []string, client peerclientcontract.Client, peerRepo peerrepositorycontract.Repository, server peerservercontract.Server) peerconnectorcontract.Connector {
	return manager{
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
		peerRepo:              peerRepo,
		server:                server,
	}
}

type manager struct {
	existingPeerAddresses []string
	client                peerclientcontract.Client
	peerRepo              peerrepositorycontract.Repository
	server                peerservercontract.Server
}

var _ peerconnectorcontract.Connector = manager{}
