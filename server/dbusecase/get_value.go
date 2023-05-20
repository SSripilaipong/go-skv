package dbusecase

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
}

type GetValueFunc func(*GetValueRequest) (*GetValueResponse, error)
