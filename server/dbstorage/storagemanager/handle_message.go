package storagemanager

import (
	"fmt"
	"go-skv/util/goutil"
)

func (m *manager) handleMessage(raw any) {
	if message, isSetMessage := raw.(SetValueMessage); isSetMessage {
		m.handleSetValueMessage(message)
	} else if message, isGetMessage := raw.(GetValueMessage); isGetMessage {
		m.handleGetValueMessage(message)
	}
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
		record = m.recordFactory.New(m.ctx)
	}
	goutil.PanicUnhandledError(record.SetValue(message))
	m.records[message.Key()] = record
}
