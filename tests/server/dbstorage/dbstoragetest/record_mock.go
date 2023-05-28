package dbstoragetest

import "go-skv/server/dbstorage"

type RecordMock struct {
	SetValue_message dbstorage.SetValueMessage
	GetValue_message dbstorage.GetValueMessage
}

func (r *RecordMock) SetValue(message dbstorage.SetValueMessage) error {
	r.SetValue_message = message
	return nil
}

func (r *RecordMock) GetValue(message dbstorage.GetValueMessage) error {
	r.GetValue_message = message
	return nil
}

func (r *RecordMock) Destroy() error {
	return nil
}
