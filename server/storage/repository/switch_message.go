package repository

import (
	. "go-skv/server/storage/repository/message"
)

func switchMessage(terminate func(Terminate), saveRecord func(SaveRecord), forwardToRecord func(msg ForwardToRecord)) func(raw any) (isTerminated bool) {
	return func(raw any) (isTerminated bool) {
		switch msg := raw.(type) {
		case Terminate:
			terminate(msg)
			return true
		case SaveRecord:
			saveRecord(msg)
		case ForwardToRecord:
			forwardToRecord(msg)
		}
		return false
	}
}
