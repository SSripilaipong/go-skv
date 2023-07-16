package dbstoragecontract

import (
	"context"
)

type Storage interface {
	Start(ctx context.Context) error
	Join() error
	GetRecord(ctx context.Context, key string, execute func(Record), failure func(err error)) error
	GetOrCreateRecord(ctx context.Context, key string, success func(Record)) error
	Add(ctx context.Context, key string, record Record) error
}
