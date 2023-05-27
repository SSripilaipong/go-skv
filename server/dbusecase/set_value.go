package dbusecase

import "go-skv/server/storage"

type SetValueRequest struct {
	Key   string
	Value string
}

type SetValueResponse struct {
}

type SetValueFunc func(*SetValueRequest) (*SetValueResponse, error)

func SetValueUsecase(dep *Dependency) SetValueFunc {
	return func(*SetValueRequest) (*SetValueResponse, error) {
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
