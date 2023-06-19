package peerconnector

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func New(existingPeerAddresses []string, client peerclientcontract.Client) peerconnectorcontract.Connector {
	return connector{
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
	}
}

type connector struct {
	existingPeerAddresses []string
	client                peerclientcontract.Client
}

var _ peerconnectorcontract.Connector = connector{}
