package dbstoragetest

import "go-skv/server/dbstorage"

type GetValueMessage struct {
	KeyField string
}

func (m *GetValueMessage) Key() string {
	return m.KeyField
}

func (m *GetValueMessage) Completed(dbstorage.GetValueResponse) error {
	return nil
}
