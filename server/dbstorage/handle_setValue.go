package dbstorage

import "go-skv/goutil"

func (s *storage) handleSetValueMessage(message SetValueMessage) {
	record := s.recordFactory.New()
	goutil.PanicUnhandledError(record.SetValue(message)) // TODO: handle error
	s.records[message.Key()] = record
}
