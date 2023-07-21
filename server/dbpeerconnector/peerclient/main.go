package peerclient

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
	"go-skv/server/dbpeerconnector/peerclient/peerclientmanager"
	"go-skv/server/dbpeerconnector/peerclient/peergrpcgateway"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func New(advertisedAddress string, replicaUpdaterFactory replicaupdatercontract.Factory) peerclientcontract.Client {
	return peerclientmanager.New(
		clientsidepeer.NewFactory(0, 0, nil),
		peergrpcgateway.NewConnector(advertisedAddress),
	)
}
