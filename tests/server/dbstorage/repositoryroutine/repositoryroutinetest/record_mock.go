package repositoryroutinetest

import (
	"go-skv/server/dbstorage/storagerecord"
)

type RecordMock struct {
	GetValue_message storagerecord.GetValueMessage
	SetValue_value   string
}

func (r *RecordMock) SetValue(value string, success func(response storagerecord.SetValueResponse)) error {
	r.SetValue_value = value
	return nil
}

func (r *RecordMock) GetValue(message storagerecord.GetValueMessage) error {
	r.GetValue_message = message
	return nil
}

func (r *RecordMock) Destroy() error {
	return nil
}
