package storagerecordfactory

import (
	"context"
	"go-skv/server/dbstorage/storagerecord"
)

type Interface interface {
	New(ctx context.Context) storagerecord.Interface
}
