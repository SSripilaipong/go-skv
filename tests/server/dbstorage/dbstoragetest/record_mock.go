package dbstoragetest

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type RecordMock struct {
	GetValue_success_response storagerecord.GetValueResponse
	GetValue_success_willFail bool
	GetValue_ctx              context.Context
	SetValue_value            string
	SetValue_ctx              context.Context
	SetValue_success_willFail bool
	SetValue_success_response storagerecord.SetValueResponse
}

func (r *RecordMock) SetValue(ctx context.Context, value string, success func(response storagerecord.SetValueResponse)) error {
	r.SetValue_ctx = ctx
	r.SetValue_value = value
	if !r.SetValue_success_willFail {
		go success(r.SetValue_success_response)
	}
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
