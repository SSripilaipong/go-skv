package storagemanager

import (
	"fmt"
	"go-skv/util/goutil"
)

func (m *manager) mainLoop() {
	for {
		select {
		case raw := <-m.ch:
			if message, isSetMessage := raw.(SetValueMessage); isSetMessage {
				m.handleSetValueMessage(message)
			} else if message, isGetMessage := raw.(GetValueMessage); isGetMessage {
				m.handleGetValueMessage(message)
			}
		case <-m.ctx.Done():
			goto stop
		}
	}
stop:
	m.stopped <- struct{}{}
}

func (m *manager) handleGetValueMessage(message GetValueMessage) {
	record, exists := m.records[message.Key()]
	if !exists {
		panic(fmt.Errorf("unhandled error"))
	}
	goutil.PanicUnhandledError(record.GetValue(message))
}

func (m *manager) handleSetValueMessage(message SetValueMessage) {
	record, exists := m.records[message.Key()]
	if !exists {
		record = m.recordFactory.New()
	}
	goutil.PanicUnhandledError(record.SetValue(message))
	m.records[message.Key()] = record
}
