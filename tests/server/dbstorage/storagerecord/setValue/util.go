package setValue

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func doSetValueWithValueAndSuccessFunc(record storagerecord.Interface, value string, success func(response storagerecord.SetValueResponse)) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.SetValue(ctx, value, success)
}

func doSetValue(record storagerecord.Interface) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.SetValue(ctx, "", func(storagerecord.SetValueResponse) {})
}

func doSetValueWithContext(record storagerecord.Interface, ctx context.Context) error {
	return record.SetValue(ctx, "", func(storagerecord.SetValueResponse) {})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
