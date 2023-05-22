package getValue

import (
	"go-skv/goutil"
	"go-skv/server/dbusecase"
)

type getValueUsecaseMock struct {
	Request *dbusecase.GetValueRequest
	Return  *dbusecase.GetValueResponse
}

func (m *getValueUsecaseMock) New() dbusecase.GetValueFunc {
	response := goutil.Coalesce(m.Return, &dbusecase.GetValueResponse{Value: ""})

	return func(request *dbusecase.GetValueRequest) (*dbusecase.GetValueResponse, error) {
		m.Request = request
		return response, nil
	}
}
