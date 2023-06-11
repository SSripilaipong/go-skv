package setValue

import (
	"context"
	"go-skv/server/dbusecase"
)

type setValueUsecaseMock struct {
	Request dbusecase.SetValueRequest
	Context context.Context
}

func (m *setValueUsecaseMock) New() dbusecase.SetValueFunc {
	return func(ctx context.Context, request dbusecase.SetValueRequest) (dbusecase.SetValueResponse, error) {
		m.Context = ctx
		m.Request = request
		return dbusecase.SetValueResponse{}, nil
	}
}
