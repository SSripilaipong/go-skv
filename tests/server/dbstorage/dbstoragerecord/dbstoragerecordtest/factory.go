package dbstoragerecordtest

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/dbstoragerecord"
)

func NewFactory() dbstorage.RecordFactory {
	return dbstoragerecord.NewFactory(context.Background(), 1)
}

func NewFactoryWithContext(ctx context.Context) dbstorage.RecordFactory {
	return dbstoragerecord.NewFactory(ctx, 1)
}
