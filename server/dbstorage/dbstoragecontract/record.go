package dbstoragecontract

import (
	"context"
)

type Record interface {
	SetValue(ctx context.Context, value string, success func(response RecordData)) error
	GetValue(ctx context.Context, success func(response RecordData)) error
	Destroy() error
}

type Factory interface {
	New(ctx context.Context) Record
}

type RecordData struct {
	Value string
}
