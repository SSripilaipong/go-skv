package setValue

import (
	"context"
	"go-skv/server/dbusecase"
)

func getStorageChannelAfterExecute(ctx context.Context, request dbusecase.SetValueRequest) chan any {
	storageChan := make(chan any, 2)
	execute := dbusecase.SetValueUsecase(dbusecase.NewDependency(storageChan))

	go func() {
		_, _ = execute(ctx, request)
	}()

	return storageChan
}
