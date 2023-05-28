package dbstorage

import (
	"go-skv/util/goutil"
)

func (s *storage) handleGetValueMessage(message GetValueMessage) {
	goutil.PanicUnhandledError(s.records[message.Key()].GetValue(message))
}
