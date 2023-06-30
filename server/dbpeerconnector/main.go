package dbpeerconnector

import (
	"go-skv/server/dbpeerconnector/connectormanager"
	"go-skv/server/dbpeerconnector/peerclient"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerserver"
)

func New(port int) peerconnectorcontract.Connector {
	return connectormanager.New([]string{"localhost:7777"}, peerclient.New(), nil, peerserver.New(port))
}
