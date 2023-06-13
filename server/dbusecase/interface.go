package dbusecase

import (
	"context"
)

type Interface interface {
	GetValue(context.Context, GetValueRequest) (GetValueResponse, error)
	SetValue(context.Context, SetValueRequest) (SetValueResponse, error)
}
