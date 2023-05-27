package dbusecase

import (
	"context"
	"fmt"
	"go-skv/server/storage"
	"time"
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
		resultChan := make(chan storage.SetValueResponse)
		dep.storageChan <- setValueMessage{key: request.Key, value: request.Value, resultChan: resultChan}
		select {
		case <-resultChan:
			return &SetValueResponse{}, nil
		case <-ctx.Done():
			return nil, fmt.Errorf("context closed")
		case <-time.After(time.Second): // TODO: parameterize
			panic(fmt.Errorf("unhandled error"))
		}
	}
}

type setValueMessage struct {
	key        string
	value      string
	resultChan chan storage.SetValueResponse
}

func (m setValueMessage) Key() string {
	return m.key
}

func (m setValueMessage) Value() string {
	return m.value
}

func (m setValueMessage) Completed(result storage.SetValueResponse) error {
	m.resultChan <- result
	return nil
}
