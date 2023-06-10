package getValue

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func doGetValueWithSuccessFunc(record storagerecord.Interface, success func(response storagerecord.GetValueResponse)) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.GetValue(ctx, success)
}

func doGetValue(record storagerecord.Interface) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.GetValue(ctx, func(storagerecord.GetValueResponse) {})
}

func doGetValueWithContext(record storagerecord.Interface, ctx context.Context) error {
	return record.GetValue(ctx, func(storagerecord.GetValueResponse) {})
}

func doSetValueWithValue(record storagerecord.Interface, value string) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.SetValue(ctx, value, func(storagerecord.SetValueResponse) {})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
