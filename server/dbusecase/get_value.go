package dbusecase

import (
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

type GetValueFunc func(*GetValueRequest) (*GetValueResponse, error)

func GetValueUsecase(dep *Dependency) GetValueFunc {
	return func(request *GetValueRequest) (*GetValueResponse, error) {
		resultChan := make(chan storage.GetValueResponse)
		dep.storageChan <- getValueMessage{key: request.Key, resultChan: resultChan}
		select {
		case result := <-resultChan:
			return &GetValueResponse{Value: result.Value}, nil
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
