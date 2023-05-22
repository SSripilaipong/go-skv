package dbusecase

import "go-skv/server/storage"

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
	Value string
}

type GetValueFunc func(*GetValueRequest) (*GetValueResponse, error)

func GetValueUsecase(dep *Dependency) GetValueFunc {
	return func(request *GetValueRequest) (*GetValueResponse, error) {
		dep.storageChan <- &storagePacket{message: storage.GetValueMessage{Key: request.Key}}
		return nil, nil
	}
}

type storagePacket struct {
	message any
}

func (s *storagePacket) Message() any {
	return s.message
}
