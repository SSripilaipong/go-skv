package dbusecase

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
	Value string
}

type GetValueFunc func(*GetValueRequest) (*GetValueResponse, error)

func GetValueUsecase(dep *Dependency) GetValueFunc {
	return func(request *GetValueRequest) (*GetValueResponse, error) {
		dep.storageChan <- &getValueMessage{key: request.Key}
		return nil, nil
	}
}

type getValueMessage struct {
	key string
}

func (m *getValueMessage) Key() string {
	return m.key
}
