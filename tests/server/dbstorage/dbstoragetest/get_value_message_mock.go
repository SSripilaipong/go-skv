package dbstoragetest

import "go-skv/server/dbstorage"

type GetValueMessage struct {
	KeyField           string
	Completed_Response dbstorage.GetValueResponse
}

func (m *GetValueMessage) Key() string {
	return m.KeyField
}

func (m *GetValueMessage) Completed(response dbstorage.GetValueResponse) error {
	m.Completed_Response = response
	return nil
}
