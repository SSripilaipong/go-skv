package dbusecase

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/util/goutil"
)

func SetValueUsecaseV2(dep Dependency) SetValueFunc {
	return func(ctx context.Context, request SetValueRequest) (SetValueResponse, error) {
		finish := make(chan struct{})
		goutil.PanicUnhandledError(dep.repo.GetOrCreateRecord(ctx, request.Key, func(record dbstorage.Record) {
			goutil.PanicUnhandledError(record.SetValue(context.Background(), request.Value, nil))
			finish <- struct{}{}
		}))
		<-finish

		return SetValueResponse{}, nil
	}
}
