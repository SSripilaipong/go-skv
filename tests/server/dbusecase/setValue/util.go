package setValue

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbserver/dbusecase"
)

func doExecuteWithRequest(usecase dbusecase.Interface, request dbusecase.SetValueRequest) (dbusecase.SetValueResponse, error) {
	return usecase.SetValue(context.Background(), request)
}

func doExecuteWithContext(usecase dbusecase.Interface, ctx context.Context) (dbusecase.SetValueResponse, error) {
	return usecase.SetValue(ctx, dbusecase.SetValueRequest{})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return goutil.NewContextWithTimeout(defaultTimeout)
}
