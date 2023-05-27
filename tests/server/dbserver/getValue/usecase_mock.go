package getValue

import (
	"context"
	"go-skv/goutil"
	"go-skv/server/dbusecase"
)

type getValueUsecaseMock struct {
	Request *dbusecase.GetValueRequest
	Context context.Context
	Return  *dbusecase.GetValueResponse
}

func (m *getValueUsecaseMock) New() dbusecase.GetValueFunc {
	response := goutil.Coalesce(m.Return, &dbusecase.GetValueResponse{Value: goutil.Pointer("")})

	return func(ctx context.Context, request *dbusecase.GetValueRequest) (*dbusecase.GetValueResponse, error) {
		m.Context = ctx
		m.Request = request
		return response, nil
	}
}
