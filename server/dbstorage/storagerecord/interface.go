package storagerecord

import (
	"context"
)

type Interface interface {
	SetValue(value string, success func(response SetValueResponse)) error
	GetValue(GetValueMessage) error
	Destroy() error
}

type Factory interface {
	New(ctx context.Context) Interface
}
