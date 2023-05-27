package dbstorage

import "go-skv/goutil"

func (s *storage) handleSetValueMessage(message SetValueMessage) {
	record := s.recordFactory.New()
	goutil.PanicUnhandledError(record.SetValue(message.Value())) // TODO: handle error
}
