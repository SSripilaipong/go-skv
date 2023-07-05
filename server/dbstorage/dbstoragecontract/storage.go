package dbstoragecontract

import (
	"context"
)

type Storage interface {
	Start(ctx context.Context) error
	Join() error
	GetRecord(ctx context.Context, key string, execute func(Record)) error
	GetOrCreateRecord(ctx context.Context, key string, success func(Record)) error
}
