package dbusecase

import (
	"context"
	"fmt"
	"go-skv/server/storage"
	"time"
)

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
	Value string
}

type GetValueFunc func(context.Context, *GetValueRequest) (*GetValueResponse, error)

func GetValueUsecase(dep *Dependency) GetValueFunc {
	return func(ctx context.Context, request *GetValueRequest) (*GetValueResponse, error) {
		resultChan := make(chan storage.GetValueResponse)
		dep.storageChan <- getValueMessage{key: request.Key, resultChan: resultChan}
		select {
		case result := <-resultChan:
			return &GetValueResponse{Value: result.Value}, nil
		case <-ctx.Done():
			return nil, fmt.Errorf("context closed")
		case <-time.After(time.Second): // TODO: parameterize
			panic(fmt.Errorf("unhandled error"))
		}
	}
}

type getValueMessage struct {
	key        string
	resultChan chan storage.GetValueResponse
}

func (m getValueMessage) Key() string {
	return m.key
}

func (m getValueMessage) Completed(result storage.GetValueResponse) error {
	m.resultChan <- result
	return nil
}
