package storagerepositorytest

import (
	"go-skv/server/dbstorage/storagerecord"
)

type GetValueMessage struct {
	KeyField           string
	Completed_Response storagerecord.GetValueResponse
}

func (m *GetValueMessage) Key() string {
	return m.KeyField
}

func (m *GetValueMessage) Completed(response storagerecord.GetValueResponse) error {
	m.Completed_Response = response
	return nil
}
