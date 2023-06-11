package setValue

import (
	"context"
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
)

func newUsecaseWithRepo(repo dbstorage.RepositoryInteractor) dbusecase.SetValueFunc {
	return dbusecase.SetValueUsecaseV2(dbusecase.NewDependency(nil, repo))
}

func doExecuteWithRequest(usecase dbusecase.SetValueFunc, request dbusecase.SetValueRequest) (dbusecase.SetValueResponse, error) {
	return usecase(context.Background(), request)
}

func doExecuteWithContext(usecase dbusecase.SetValueFunc, ctx context.Context) (dbusecase.SetValueResponse, error) {
	return usecase(ctx, dbusecase.SetValueRequest{})
}
