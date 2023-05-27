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
		dep.storageChan <- setValueMessage{key: request.Key}
		return nil, nil
	}
}

type setValueMessage struct {
	key string
}

func (m setValueMessage) Key() string {
	return m.key
}

func (m setValueMessage) Value() string {
	return ""
}

func (m setValueMessage) Completed(storage.SetValueResponse) error {
	return nil
}
