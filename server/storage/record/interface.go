package record

import "context"

type Factory interface {
	New(ctx context.Context, value string) chan<- any
	NewReplica(ctx context.Context, value string) chan<- any
}
