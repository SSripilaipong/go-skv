package dbusecase

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
	Value string
}

type GetValueFunc func(*GetValueRequest) (*GetValueResponse, error)
