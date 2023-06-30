package peerclientmanagertest

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerclient/peerclientmanager"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway/peergrpcgatewaycontract"
)

type dependency struct {
	peerFactory      clientsidepeercontract.Factory
	gatewayConnector peergrpcgatewaycontract.GatewayConnector
}

func defaultDependency() dependency {
	return dependency{
		peerFactory: &PeerFactoryMock{},
		gatewayConnector: &GatewayConnectorMock{
			ConnectTo_Return: &GatewayMock{},
		},
	}
}

func New(options ...func(*dependency)) peerclientcontract.Client {
	deps := defaultDependency()
	for _, option := range options {
		option(&deps)
	}
	return peerclientmanager.New(deps.peerFactory, deps.gatewayConnector)
}

func WithPeerFactory(peerFactory clientsidepeercontract.Factory) func(*dependency) {
	return func(deps *dependency) {
		deps.peerFactory = peerFactory
	}
}

func WithGatewayConnector(gatewayConnector peergrpcgatewaycontract.GatewayConnector) func(*dependency) {
	return func(deps *dependency) {
		deps.gatewayConnector = gatewayConnector
	}
}
