package dbstorageTest

import "go-skv/server/dbstorage"

type SetValueMessage struct {
	KeyField   string
	ValueField string
}

func (m *SetValueMessage) Key() string {
	return m.KeyField
}

func (m *SetValueMessage) Value() string {
	return m.ValueField
}

func (m *SetValueMessage) Completed(dbstorage.SetValueResponse) error {
	return nil
}
