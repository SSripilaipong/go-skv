package storagerecord

import (
	"context"
)

type Interface interface {
	SetValue(SetValueMessage) error
	GetValue(GetValueMessage) error
	Destroy() error
}

type Factory interface {
	New(ctx context.Context) Interface
}
