package setValue

import "go-skv/server/dbstorage"

func NewMessage() dbstorage.SetValueMessage {
	return nil
}

type message struct {
	key   string
	value string
}

func (m *message) Key() string {
	return m.key
}

func (m *message) Value() string {
	return m.value
}

func (m *message) Completed(dbstorage.SetValueResponse) error {
	return nil
}
