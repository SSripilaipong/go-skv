package repositoryinteractor

import (
	"context"
	"go-skv/server/dbstorage/repositoryroutine"
)

type Interface interface {
	GetRecord(ctx context.Context, key string, success repositoryroutine.GetRecordSuccessCallback) error
	GetOrCreateRecord(ctx context.Context, key string, success repositoryroutine.GetOrCreateRecordSuccessCallback) error
}
