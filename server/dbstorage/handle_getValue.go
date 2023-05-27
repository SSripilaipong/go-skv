package dbstorage

import "go-skv/goutil"

func (s *storage) handleGetValueMessage(message GetValueMessage) {
	goutil.PanicUnhandledError(s.records[message.Key()].GetValue(message))
}
