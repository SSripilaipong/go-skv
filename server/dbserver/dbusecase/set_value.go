package dbusecase

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type SetValueRequest struct {
	Key   string
	Value string
}

type SetValueResponse struct {
}

func (u usecase) SetValue(ctx context.Context, request SetValueRequest) (SetValueResponse, error) {
	completed := make(chan struct{})
	doSignalCompleted := func(dbstoragecontract.RecordData) { completed <- struct{}{} }
	doSetValueToRecord := func(record dbstoragecontract.Record) {
		goutil.PanicUnhandledError(record.SetValue(ctx, request.Value, doSignalCompleted))
	}

	goutil.PanicUnhandledError(u.repo.GetOrCreateRecord(ctx, request.Key, doSetValueToRecord))

	select {
	case <-completed:
		return SetValueResponse{}, nil
	case <-ctx.Done():
		return SetValueResponse{}, ContextCancelledError{}
	}
}
