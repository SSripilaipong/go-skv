package getValue

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbusecase"
)

func doExecute(usecase dbusecase.Interface) (dbusecase.GetValueResponse, error) {
	return usecase.GetValue(context.Background(), dbusecase.GetValueRequest{})
}

func doExecuteWithRequest(usecase dbusecase.Interface, request dbusecase.GetValueRequest) (dbusecase.GetValueResponse, error) {
	return usecase.GetValue(context.Background(), request)
}

func doExecuteWithContext(usecase dbusecase.Interface, ctx context.Context) (dbusecase.GetValueResponse, error) {
	return usecase.GetValue(ctx, dbusecase.GetValueRequest{})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return goutil.NewContextWithTimeout(defaultTimeout)
}
