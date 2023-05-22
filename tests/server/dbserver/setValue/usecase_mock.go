package setValue

import "go-skv/server/dbusecase"

type setValueUsecaseMock struct {
	Request *dbusecase.SetValueRequest
}

func (m *setValueUsecaseMock) New() dbusecase.SetValueFunc {
	return func(request *dbusecase.SetValueRequest) (*dbusecase.SetValueResponse, error) {
		m.Request = request
		return nil, nil
	}
}
