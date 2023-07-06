package peerclient

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerclient/peerclientmanager"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway"
)

func New(advertisedAddress string) peerclientcontract.Client {
	return peerclientmanager.New(
		clientsidepeer.NewFactory(nil),
		peergrpcgateway.NewConnector(advertisedAddress),
	)
}
