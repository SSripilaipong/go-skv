package dbstorageTest

import "go-skv/server/dbstorage"

type RecordMock struct {
	SetValue_message dbstorage.SetValueMessage
}

func (r *RecordMock) SetValue(message dbstorage.SetValueMessage) error {
	r.SetValue_message = message
	return nil
}
