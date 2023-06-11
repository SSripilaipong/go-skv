package dbusecase

import (
	"context"
	"go-skv/util/goutil"
)

func SetValueUsecaseV2(dep Dependency) SetValueFunc {
	return func(ctx context.Context, request SetValueRequest) (SetValueResponse, error) {
		goutil.PanicUnhandledError(dep.repo.GetOrCreateRecord(ctx, request.Key, nil))
		return SetValueResponse{}, nil
	}
}
