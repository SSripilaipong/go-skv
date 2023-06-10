package getValue

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

func doGetValueWithSuccessFunc(record storagerecord.Interface, success func(response storagerecord.GetValueResponse)) error {
	ctx, _ := contextWithDefaultTimeout()
	return record.GetValue(ctx, success)
}

func doGetValueWithContext(record storagerecord.Interface, ctx context.Context) error {
	return record.GetValue(ctx, func(storagerecord.GetValueResponse) {})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
