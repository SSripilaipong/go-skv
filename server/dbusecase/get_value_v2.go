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
		goutil.PanicUnhandledError(dep.repo.GetRecord(ctx, request.Key, func(record dbstorage.Record) {
			goutil.PanicUnhandledError(record.GetValue(nil, func(response dbstorage.GetValueResponse) {
				resultCh <- response
			}))
		}))

		result := <-resultCh
		return GetValueResponse{Value: result.Value}, nil
	}
}
