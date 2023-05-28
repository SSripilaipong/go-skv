package dbstoragetest

import "go-skv/server/dbstorage"

type SetValueMessage struct {
	KeyField           string
	ValueField         string
	Completed_Response *dbstorage.SetValueResponse
}

func (m *SetValueMessage) Key() string {
	return m.KeyField
}

func (m *SetValueMessage) Value() string {
	return m.ValueField
}

func (m *SetValueMessage) Completed(response dbstorage.SetValueResponse) error {
	m.Completed_Response = &response
	return nil
}
