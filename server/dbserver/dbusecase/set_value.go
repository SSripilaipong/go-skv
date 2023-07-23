package dbusecase

import (
	"context"
	"fmt"
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
		fmt.Printf("DbServer: SetValue(%#v, %#v)\n", request.Key, request.Value) // TODO: remove demo log
		return SetValueResponse{}, nil
	case <-ctx.Done():
		return SetValueResponse{}, ContextCancelledError{}
	}
}
