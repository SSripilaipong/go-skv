package peergrpctest

import (
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"
)

type connectorDependency struct {
	address string
}

func defaultConnectorDependency() connectorDependency {
	return connectorDependency{
		address: "0.0.0.0:0000",
	}
}

func NewConnector(options ...func(*connectorDependency)) peergrpcgatewaycontract.GatewayConnector {
	deps := defaultConnectorDependency()
	for _, option := range options {
		option(&deps)
	}
	return peergrpcgateway.NewConnector(deps.address)
}

func WithAdvertisedAddress(address string) func(*connectorDependency) {
	return func(deps *connectorDependency) {
		deps.address = address
	}
}
