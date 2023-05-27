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
	return func(context.Context, *SetValueRequest) (*SetValueResponse, error) {
		dep.storageChan <- setValueMessage{}
		return nil, nil
	}
}

type setValueMessage struct{}

func (m setValueMessage) Key() string {
	return ""
}

func (m setValueMessage) Value() string {
	return ""
}

func (m setValueMessage) Completed(storage.SetValueResponse) error {
	return nil
}
