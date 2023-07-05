package storagerepository

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type Interface interface {
	Start(ctx context.Context) error
	Join() error
	GetRecord(ctx context.Context, key string, execute func(storagerecord.Interface)) error
	GetOrCreateRecord(ctx context.Context, key string, success func(storagerecord.Interface)) error
}

type Interactor = Interface
