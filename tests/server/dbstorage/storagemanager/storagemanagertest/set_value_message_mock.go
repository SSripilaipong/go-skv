package storagemanagertest

import (
	"go-skv/server/dbstorage/storagemanager"
)

type SetValueMessage struct {
	KeyField           string
	ValueField         string
	Completed_Response *storagemanager.SetValueResponse
}

func (m *SetValueMessage) Key() string {
	return m.KeyField
}

func (m *SetValueMessage) Value() string {
	return m.ValueField
}

func (m *SetValueMessage) Completed(response storagemanager.SetValueResponse) error {
	m.Completed_Response = &response
	return nil
}
