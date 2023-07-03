package dbusecase

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage"
)

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
	Value string
}

func (u usecase) GetValue(ctx context.Context, request GetValueRequest) (GetValueResponse, error) {
	resultCh := make(chan dbstorage.GetValueResponse)
	doSendResultBack := func(response dbstorage.GetValueResponse) { resultCh <- response }
	doReadRecordThenSendResultBack := func(record dbstorage.Record) {
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
