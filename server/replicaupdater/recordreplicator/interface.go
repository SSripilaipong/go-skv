package recordreplicator

import (
	"context"
)

type Factory interface {
	New(ctx context.Context, key string, value string) chan<- any
}
