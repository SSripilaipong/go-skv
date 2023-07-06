package clientsidepeertest

import "go-skv/server/replicaupdater/replicaupdatercontract"

type ReplicaInboundUpdaterMock struct {
	Update_key   string
	Update_value string
}

var _ replicaupdatercontract.InboundUpdater = &ReplicaInboundUpdaterMock{}

func (u *ReplicaInboundUpdaterMock) Update(key string, value string) error {
	u.Update_key = key
	u.Update_value = value

	return nil
}
