package dbpeerconnector

import (
	"go-skv/server/dbpeerconnector/connectormanager"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func New() peerconnectorcontract.Connector {
	return connectormanager.New([]string{"localhost:5555"}, nil, nil, nil)
}
