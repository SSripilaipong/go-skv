package repositoryroutinetest

import (
	"go-skv/server/dbstorage/storagerecord"
)

type RecordMock struct {
	SetValue_message storagerecord.SetValueMessage
	GetValue_message storagerecord.GetValueMessage
}

func (r *RecordMock) SetValue(message storagerecord.SetValueMessage) error {
	r.SetValue_message = message
	return nil
}

func (r *RecordMock) GetValue(message storagerecord.GetValueMessage) error {
	r.GetValue_message = message
	return nil
}

func (r *RecordMock) Destroy() error {
	return nil
}
