package dbusecase

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/util/goutil"
)

func SetValueUsecaseV2(dep Dependency) SetValueFunc {
	return func(ctx context.Context, request SetValueRequest) (SetValueResponse, error) {
		finish := make(chan struct{})
		goutil.PanicUnhandledError(dep.repo.GetOrCreateRecord(ctx, request.Key, func(record dbstorage.Record) {
			goutil.PanicUnhandledError(record.SetValue(ctx, request.Value, func(storagerecord.SetValueResponse) {
				finish <- struct{}{}
			}))
		}))

		select {
		case <-finish:
			return SetValueResponse{}, nil
		case <-ctx.Done():
			return SetValueResponse{}, ContextCancelledError{}
		}
	}
}
