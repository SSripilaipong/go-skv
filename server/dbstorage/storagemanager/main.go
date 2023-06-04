package storagemanager

import "context"

func New(ctx context.Context, ch chan any, recordFactory RecordFactory) Interface {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	return &manager{
		ch:            ch,
		recordFactory: recordFactory,
		ctx:           ctxWithCancel,
		cancel:        cancel,

		stopped: make(chan struct{}),
		records: make(map[string]DbRecord),
	}
}
