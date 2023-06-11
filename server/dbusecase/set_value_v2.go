package dbusecase

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/util/goutil"
)

func SetValueUsecaseV2(dep Dependency) SetValueFunc {
	return func(ctx context.Context, request SetValueRequest) (SetValueResponse, error) {
		completed := make(chan struct{})
		doSignalCompleted := func(storagerecord.SetValueResponse) { completed <- struct{}{} }
		doSetValueToRecord := func(record dbstorage.Record) {
			goutil.PanicUnhandledError(record.SetValue(ctx, request.Value, doSignalCompleted))
		}

		goutil.PanicUnhandledError(dep.repo.GetOrCreateRecord(ctx, request.Key, doSetValueToRecord))

		select {
		case <-completed:
			return SetValueResponse{}, nil
		case <-ctx.Done():
			return SetValueResponse{}, ContextCancelledError{}
		}
	}
}
