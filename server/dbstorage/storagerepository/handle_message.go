package storagerepository

import (
	"fmt"
)

func (m *manager) handleMessage(raw any) {
	switch message := raw.(type) {
	case GetOrCreateRecordMessage:
		m.handleGetOrCreateRecord(message)
	case GetRecordMessage:
		m.handleGetRecord(message)
	}
}

func (m *manager) handleGetOrCreateRecord(message GetOrCreateRecordMessage) {
	record, exists := m.records[message.Key]
	if !exists {
		record = m.recordFactory.New(m.ctx)
	}
	message.Success(record)
	m.records[message.Key] = record
}

func (m *manager) handleGetRecord(message GetRecordMessage) {
	record, exists := m.records[message.Key]
	if !exists {
		panic(fmt.Errorf("unhandled error"))
	}
	message.Success(record)
}
