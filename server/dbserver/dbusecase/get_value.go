package dbusecase

import (
	"context"
	"fmt"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
	Value string
}

func (u usecase) GetValue(ctx context.Context, request GetValueRequest) (GetValueResponse, error) {
	resultCh := make(chan dbstoragecontract.RecordData)
	doSendResultBack := func(response dbstoragecontract.RecordData) { resultCh <- response }
	doReadRecordThenSendResultBack := func(record dbstoragecontract.Record) {
		goutil.PanicUnhandledError(record.GetValue(ctx, doSendResultBack))
	}

	goutil.PanicUnhandledError(u.repo.GetRecord(ctx, request.Key, doReadRecordThenSendResultBack, func(error) {}))

	select {
	case result := <-resultCh:
		fmt.Printf("DbServer: GetValue(%#v) -> %#v\n", request.Key, result.Value) // TODO: remove demo log
		return GetValueResponse{Value: result.Value}, nil
	case <-ctx.Done():
		return GetValueResponse{}, ContextCancelledError{}
	}
}
