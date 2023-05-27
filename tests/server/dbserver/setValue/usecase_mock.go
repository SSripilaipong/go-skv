package setValue

import (
	"context"
	"go-skv/server/dbusecase"
)

type setValueUsecaseMock struct {
	Request *dbusecase.SetValueRequest
}

func (m *setValueUsecaseMock) New() dbusecase.SetValueFunc {
	return func(ctx context.Context, request *dbusecase.SetValueRequest) (*dbusecase.SetValueResponse, error) {
		m.Request = request
		return nil, nil
	}
}
