package setValue

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func doSetValueWithValueAndSuccessFunc(record dbstoragecontract.Record, value string, success func(response dbstoragecontract.RecordData)) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.SetValue(ctx, value, success)
}

func doSetValue(record dbstoragecontract.Record) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.SetValue(ctx, "", func(dbstoragecontract.RecordData) {})
}

func doSetValueWithContext(record dbstoragecontract.Record, ctx context.Context) error {
	return record.SetValue(ctx, "", func(dbstoragecontract.RecordData) {})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
