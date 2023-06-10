package storagerecord

import (
	"context"
)

type Interface interface {
	SetValue(ctx context.Context, value string, success func(response SetValueResponse)) error
	GetValue(ctx context.Context, success func(response GetValueResponse)) error
	Destroy() error
}

type Factory interface {
	New(ctx context.Context) Interface
}
