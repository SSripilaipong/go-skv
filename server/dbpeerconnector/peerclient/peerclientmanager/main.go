package peerclientmanager

import (
	"go-skv/server/dbpeerconnector/peerclient/clientsidepeer/clientsidepeercontract"
	"go-skv/server/dbpeerconnector/peerclient/peerclientcontract"
)

func New(peerFactory clientsidepeercontract.Factory) peerclientcontract.Client {
	return client{
		peerFactory: peerFactory,
	}
}

type client struct {
	peerFactory clientsidepeercontract.Factory
}
