package peerconnector

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func New(existingPeerAddresses []string, client peerclientcontract.Client, listener peerconnectorcontract.UpdateListener) peerconnectorcontract.Connector {
	return connector{
		existingPeerAddresses: existingPeerAddresses,
		client:                client,
		listener:              listener,
	}
}

type connector struct {
	existingPeerAddresses []string
	client                peerclientcontract.Client
	listener              peerconnectorcontract.UpdateListener
}

var _ peerconnectorcontract.Connector = connector{}
