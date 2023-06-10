package dbusecase

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/util/goutil"
)

type GetValueFuncV2 func(context.Context, GetValueRequest) (GetValueResponse, error)

func GetValueUsecaseV2(dep Dependency) GetValueFuncV2 {
	return func(ctx context.Context, request GetValueRequest) (GetValueResponse, error) {
		resultCh := make(chan dbstorage.GetValueResponse)
		doSendResultBack := func(response dbstorage.GetValueResponse) { resultCh <- response }
		doReadRecordThenSendResultBack := func(record dbstorage.Record) {
			goutil.PanicUnhandledError(record.GetValue(ctx, doSendResultBack))
		}

		goutil.PanicUnhandledError(dep.repo.GetRecord(ctx, request.Key, doReadRecordThenSendResultBack))

		select {
		case result := <-resultCh:
			return GetValueResponse{Value: result.Value}, nil
		case <-ctx.Done():
			return GetValueResponse{}, ContextCancelledError{}
		}
	}
}
