package dbstorage

import "go-skv/goutil"

func (s *storage) handleSetValueMessage(message SetValueMessage) {
	record, exists := s.records[message.Key()]
	if !exists {
		record = s.recordFactory.New()
	}
	goutil.PanicUnhandledError(record.SetValue(message)) // TODO: handle error
	s.records[message.Key()] = record
}
