package dbusecase

type SetValueRequest struct {
	Key   string
	Value string
}

type SetValueResponse struct {
}

type SetValueFunc func(*SetValueRequest) (*SetValueResponse, error)
