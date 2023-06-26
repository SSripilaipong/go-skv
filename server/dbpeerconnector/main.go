package dbpeerconnector

import (
	"go-skv/server/dbpeerconnector/peerconnector"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func New() peerconnectorcontract.Connector {
	return peerconnector.New(nil, nil, nil, nil)
}
