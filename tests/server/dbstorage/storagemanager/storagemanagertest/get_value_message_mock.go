package storagemanagertest

import (
	"go-skv/server/dbstorage/storagemanager"
)

type GetValueMessage struct {
	KeyField           string
	Completed_Response storagemanager.GetValueResponse
}

func (m *GetValueMessage) Key() string {
	return m.KeyField
}

func (m *GetValueMessage) Completed(response storagemanager.GetValueResponse) error {
	m.Completed_Response = response
	return nil
}
