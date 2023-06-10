package dbusecase

import (
	"context"
	"go-skv/util/goutil"
)

type GetValueFuncV2 func(context.Context, GetValueRequest) (GetValueResponse, error)

func GetValueUsecaseV2(dep Dependency) GetValueFuncV2 {
	return func(ctx context.Context, request GetValueRequest) (GetValueResponse, error) {
		goutil.PanicUnhandledError(dep.repo.GetRecord(ctx, request.Key, nil))

		return GetValueResponse{}, nil
	}
}
