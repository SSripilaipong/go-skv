package storagerepository

import "context"

type Interface interface {
	Start() error
	Stop() error
}

type Interactor interface {
	GetRecord(ctx context.Context, key string, success GetRecordSuccessCallback) error
	GetOrCreateRecord(ctx context.Context, key string, success GetOrCreateRecordSuccessCallback) error
}
