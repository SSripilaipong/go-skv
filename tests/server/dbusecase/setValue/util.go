package setValue

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
	"go-skv/util/goutil"
)

func newUsecaseWithRepo(repo dbstorage.RepositoryInteractor) dbusecase.SetValueFunc {
	return dbusecase.SetValueUsecase(dbusecase.NewDependency(nil, repo))
}

func doExecuteWithRequest(usecase dbusecase.SetValueFunc, request dbusecase.SetValueRequest) (dbusecase.SetValueResponse, error) {
	return usecase(context.Background(), request)
}

func doExecuteWithContext(usecase dbusecase.SetValueFunc, ctx context.Context) (dbusecase.SetValueResponse, error) {
	return usecase(ctx, dbusecase.SetValueRequest{})
}

func contextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return goutil.NewContextWithTimeout(defaultTimeout)
}
