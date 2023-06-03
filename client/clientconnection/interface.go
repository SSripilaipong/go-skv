package clientconnection

import "context"

type Interface interface {
	GetValue(ctx context.Context, key string) (string, error)
	SetValue(ctx context.Context, key string, value string) error
	Close() error
}
