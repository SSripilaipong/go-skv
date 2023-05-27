package dbusecase

import (
	"context"
	"go-skv/server/storage"
)

type SetValueRequest struct {
	Key   string
	Value string
}

type SetValueResponse struct {
}

type SetValueFunc func(context.Context, *SetValueRequest) (*SetValueResponse, error)

func SetValueUsecase(dep *Dependency) SetValueFunc {
	return func(ctx context.Context, request *SetValueRequest) (*SetValueResponse, error) {
		dep.storageChan <- setValueMessage{key: request.Key, value: request.Value}
		return &SetValueResponse{}, nil
	}
}

type setValueMessage struct {
	key   string
	value string
}

func (m setValueMessage) Key() string {
	return m.key
}

func (m setValueMessage) Value() string {
	return m.value
}

func (m setValueMessage) Completed(storage.SetValueResponse) error {
	return nil
}
