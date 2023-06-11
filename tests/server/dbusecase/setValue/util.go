package setValue

import (
	"context"
	"go-skv/server/dbusecase"
)

func getStorageChannelAfterExecute(ctx context.Context, request dbusecase.SetValueRequest) chan any {
	storageChan := make(chan any, 2)
	execute := newUsecaseWithStorageChan(storageChan)

	go func() {
		_, _ = execute(ctx, request)
	}()

	return storageChan
}

func newUsecase() dbusecase.SetValueFunc {
	return dbusecase.SetValueUsecase(dbusecase.NewDependency(make(chan any, 1), nil))
}

func newUsecaseWithStorageChan(storageChan chan<- any) dbusecase.SetValueFunc {
	return dbusecase.SetValueUsecase(dbusecase.NewDependency(storageChan, nil))
}

func newClosedContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}
