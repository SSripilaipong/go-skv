package getValue

import (
	"context"
	"go-skv/server/dbusecase"
)

func getStorageChannelAfterExecute(ctx context.Context, request *dbusecase.GetValueRequest) chan any {
	storageChan := make(chan any, 2)
	execute := dbusecase.GetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		_, _ = execute(ctx, request)
	}()

	return storageChan
}
