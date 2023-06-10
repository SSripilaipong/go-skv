package storagerecordtest

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func DoNewRecord(factory storagerecord.Factory) storagerecord.Interface {
	return factory.New(context.Background())
}

func DoNewRecordWithContext(factory storagerecord.Factory, ctx context.Context) storagerecord.Interface {
	return factory.New(ctx)
}
