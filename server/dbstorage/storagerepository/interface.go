package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type Interface interface {
	Start(ctx context.Context) error
	Join() error
}

type Interactor interface {
	GetRecord(ctx context.Context, key string, execute func(storagerecord.Interface)) error
	GetOrCreateRecord(ctx context.Context, key string, success GetOrCreateRecordSuccessCallback) error
}
