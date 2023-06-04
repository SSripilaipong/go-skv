package storagerecordtest

import (
	"context"
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/server/dbstorage/storagerecord"
)

func NewFactory() storagemanager.RecordFactory {
	return storagerecord.NewFactory(context.Background(), 1)
}

func NewFactoryWithContext(ctx context.Context) storagemanager.RecordFactory {
	return storagerecord.NewFactory(ctx, 1)
}
