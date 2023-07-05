package dbusecase

import (
	"context"
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

	goutil.PanicUnhandledError(u.repo.GetRecord(ctx, request.Key, doReadRecordThenSendResultBack))

	select {
	case result := <-resultCh:
		return GetValueResponse{Value: result.Value}, nil
	case <-ctx.Done():
		return GetValueResponse{}, ContextCancelledError{}
	}
}
