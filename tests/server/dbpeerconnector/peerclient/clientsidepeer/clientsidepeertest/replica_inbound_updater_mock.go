package clientsidepeertest

import (
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

type ReplicaInboundUpdaterMock struct {
	Update_key   string
	Update_value string
	Update_Do    func(key, value string)
}

var _ replicaupdatercontract.InboundUpdater = &ReplicaInboundUpdaterMock{}

func (u *ReplicaInboundUpdaterMock) Update(key string, value string) error {
	u.Update_key = key
	u.Update_value = value

	if u.Update_Do != nil {
		u.Update_Do(key, value)
	}

	return nil
}
