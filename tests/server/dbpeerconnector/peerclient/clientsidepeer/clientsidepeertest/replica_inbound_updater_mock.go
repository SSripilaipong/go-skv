package clientsidepeertest

import "go-skv/server/replicaupdater/replicaupdatercontract"

type ReplicaInboundUpdaterMock struct{}

var _ replicaupdatercontract.InboundUpdater = &ReplicaInboundUpdaterMock{}
