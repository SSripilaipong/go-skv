package peergrpcgateway

import "go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"

func NewConnector(advertisedAddress string) peergrpcgatewaycontract.GatewayConnector {
	return connector{advertisedAddress: advertisedAddress}
}
