package dbstoragetest

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type RecordMock struct {
	GetValue_message          storagerecord.GetValueMessage
	SetValue_value            string
	GetValue_success_response storagerecord.GetValueResponse
}

func (r *RecordMock) SetValue(value string, success func(response storagerecord.SetValueResponse)) error {
	r.SetValue_value = value
	return nil
}

func (r *RecordMock) GetValue(ctx context.Context, success func(response storagerecord.GetValueResponse)) error {
	go success(r.GetValue_success_response)
	return nil
}

func (r *RecordMock) Destroy() error {
	return nil
}
