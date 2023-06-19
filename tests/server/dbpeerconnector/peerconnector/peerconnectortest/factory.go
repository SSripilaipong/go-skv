package peerconnectortest

import (
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerconnector"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
)

func NewWithAddressesAndClient(addresses []string, client peerclientcontract.Client) peerconnectorcontract.Connector {
	return peerconnector.New(addresses, client, &UpdateListenerMock{})
}

func NewWithClientAndUpdateListener(client peerclientcontract.Client, listener peerconnectorcontract.UpdateListener) peerconnectorcontract.Connector {
	return peerconnector.New([]string{"0.0.0.0:9999"}, client, listener)
}
