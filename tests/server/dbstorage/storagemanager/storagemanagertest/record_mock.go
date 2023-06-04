package storagemanagertest

import (
	"go-skv/server/dbstorage/storagemanager"
)

type RecordMock struct {
	SetValue_message storagemanager.SetValueMessage
	GetValue_message storagemanager.GetValueMessage
}

func (r *RecordMock) SetValue(message storagemanager.SetValueMessage) error {
	r.SetValue_message = message
	return nil
}

func (r *RecordMock) GetValue(message storagemanager.GetValueMessage) error {
	r.GetValue_message = message
	return nil
}

func (r *RecordMock) Destroy() error {
	return nil
}
