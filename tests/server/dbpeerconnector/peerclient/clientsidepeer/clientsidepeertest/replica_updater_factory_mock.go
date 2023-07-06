package clientsidepeertest

import "go-skv/server/replicaupdater/replicaupdatercontract"

type ReplicaUpdaterFactoryMock struct {
	NewInboundUpdater_IsCalled bool
}

func (r *ReplicaUpdaterFactoryMock) NewInboundUpdater() (replicaupdatercontract.InboundUpdater, error) {
	r.NewInboundUpdater_IsCalled = true
	return nil, nil
}

var _ replicaupdatercontract.Factory = &ReplicaUpdaterFactoryMock{}
