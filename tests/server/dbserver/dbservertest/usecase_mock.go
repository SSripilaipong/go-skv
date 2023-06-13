package dbservertest

import (
	"context"
	"go-skv/server/dbusecase"
)

type UsecaseMock struct {
	GetValue_Context context.Context
	GetValue_Request dbusecase.GetValueRequest
	GetValue_Return  dbusecase.GetValueResponse
	SetValue_Request dbusecase.SetValueRequest
	SetValue_Context context.Context
}

func (u *UsecaseMock) GetValue(ctx context.Context, request dbusecase.GetValueRequest) (dbusecase.GetValueResponse, error) {
	u.GetValue_Context = ctx
	u.GetValue_Request = request
	return u.GetValue_Return, nil
}

func (u *UsecaseMock) SetValue(ctx context.Context, request dbusecase.SetValueRequest) (dbusecase.SetValueResponse, error) {
	u.SetValue_Context = ctx
	u.SetValue_Request = request
	return dbusecase.SetValueResponse{}, nil
}

var _ dbusecase.Interface = &UsecaseMock{}
