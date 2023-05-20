package getValue

import "go-skv/server/dbusecase"

type getValueUsecaseMock struct {
	Request *dbusecase.GetValueRequest
}

func (m *getValueUsecaseMock) New() dbusecase.GetValueFunc {
	return func(request *dbusecase.GetValueRequest) (*dbusecase.GetValueResponse, error) {
		m.Request = request
		return nil, nil
	}
}
