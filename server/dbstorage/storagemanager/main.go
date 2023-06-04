package storagemanager

import "context"

func New(ch chan any, recordFactory RecordFactory) Interface {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	return &manager{
		ch:            ch,
		recordFactory: recordFactory,
		ctx:           ctxWithCancel,
		cancel:        cancel,

		stopped: make(chan struct{}),
		records: make(map[string]DbRecord),
	}
}
