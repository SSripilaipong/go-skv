package storagemanagertest

import (
	"go-skv/server/dbstorage/storagerecord"
)

type SetValueMessage struct {
	KeyField           string
	ValueField         string
	Completed_Response *storagerecord.SetValueResponse
}

func (m *SetValueMessage) Key() string {
	return m.KeyField
}

func (m *SetValueMessage) Value() string {
	return m.ValueField
}

func (m *SetValueMessage) Completed(response storagerecord.SetValueResponse) error {
	m.Completed_Response = &response
	return nil
}
