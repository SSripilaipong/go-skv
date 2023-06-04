package storagemanager

import "context"

type Interface interface {
	Start() error
	Stop() error
}

type RecordFactory interface {
	New(ctx context.Context) DbRecord
}

type DbRecord interface {
	SetValue(SetValueMessage) error
	GetValue(GetValueMessage) error
	Destroy() error
}
