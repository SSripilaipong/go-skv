package storagerecordtest

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func DoNewRecord(factory dbstoragecontract.Factory) dbstoragecontract.Record {
	return factory.New(context.Background())
}

func DoNewRecordWithContext(factory dbstoragecontract.Factory, ctx context.Context) dbstoragecontract.Record {
	return factory.New(ctx)
}
