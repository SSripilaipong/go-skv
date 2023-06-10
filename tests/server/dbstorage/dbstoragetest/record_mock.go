package dbstoragetest

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type RecordMock struct {
	GetValue_message          storagerecord.GetValueMessage
	SetValue_value            string
	GetValue_success_response storagerecord.GetValueResponse
	GetValue_success_willFail bool
	GetValue_ctx              context.Context
}

func (r *RecordMock) SetValue(ctx context.Context, value string, success func(response storagerecord.SetValueResponse)) error {
	r.SetValue_value = value
	return nil
}

func (r *RecordMock) GetValue(ctx context.Context, success func(response storagerecord.GetValueResponse)) error {
	r.GetValue_ctx = ctx
	if !r.GetValue_success_willFail {
		go success(r.GetValue_success_response)
	}
	return nil
}

func (r *RecordMock) Destroy() error {
	return nil
}
