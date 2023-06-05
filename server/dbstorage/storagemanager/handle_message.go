package storagemanager

import (
	"fmt"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/util/goutil"
)

func (m *manager) handleMessage(raw any) {
	if message, isSetMessage := raw.(storagerecord.SetValueMessage); isSetMessage {
		m.handleSetValueMessage(message)
	} else if message, isGetMessage := raw.(storagerecord.GetValueMessage); isGetMessage {
		m.handleGetValueMessage(message)
	}
}

func (m *manager) handleGetValueMessage(message storagerecord.GetValueMessage) {
	record, exists := m.records[message.Key()]
	if !exists {
		panic(fmt.Errorf("unhandled error"))
	}
	goutil.PanicUnhandledError(record.GetValue(message))
}

func (m *manager) handleSetValueMessage(message storagerecord.SetValueMessage) {
	record, exists := m.records[message.Key()]
	if !exists {
		record = m.recordFactory.New(m.ctx)
	}
	goutil.PanicUnhandledError(record.SetValue(message))
	m.records[message.Key()] = record
}
