package dbpeerconnector

import (
	"go-skv/server/dbpeerconnector/connectormanager"
	"go-skv/server/dbpeerconnector/peerclient"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerserver"
)

func New(port int, advertisedAddress string, existingPeerAddresses []string) peerconnectorcontract.Connector {
	return connectormanager.New(
		existingPeerAddresses,
		peerclient.New(advertisedAddress),
		nil,
		peerserver.New(port),
	)
}
