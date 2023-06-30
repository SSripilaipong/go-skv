package peerclientmanager

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"
)

func New(peerFactory clientsidepeercontract.Factory, gatewayConnector peergrpcgatewaycontract.GatewayConnector) peerclientcontract.Client {
	return client{
		peerFactory:      peerFactory,
		gatewayConnector: gatewayConnector,
	}
}

type client struct {
	peerFactory      clientsidepeercontract.Factory
	gatewayConnector peergrpcgatewaycontract.GatewayConnector
}
