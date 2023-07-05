package getValue

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func doGetValueWithSuccessFunc(record dbstoragecontract.Record, success func(response dbstoragecontract.RecordData)) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.GetValue(ctx, success)
}

func doGetValue(record dbstoragecontract.Record) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.GetValue(ctx, func(dbstoragecontract.RecordData) {})
}

func doGetValueWithContext(record dbstoragecontract.Record, ctx context.Context) error {
	return record.GetValue(ctx, func(dbstoragecontract.RecordData) {})
}

func doSetValueWithValue(record dbstoragecontract.Record, value string) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.SetValue(ctx, value, func(dbstoragecontract.RecordData) {})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
