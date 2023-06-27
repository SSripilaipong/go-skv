package dbpeerconnector

import (
	"go-skv/server/dbpeerconnector/connectormanager"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func New() peerconnectorcontract.Connector {
	return connectormanager.New(nil, nil, nil)
}
