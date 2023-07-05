package storagerepository

import "context"

type Interface interface {
	Start(ctx context.Context) error
	Join() error
}

type Interactor interface {
	GetRecord(ctx context.Context, key string, success GetRecordSuccessCallback) error
	GetOrCreateRecord(ctx context.Context, key string, success GetOrCreateRecordSuccessCallback) error
}
