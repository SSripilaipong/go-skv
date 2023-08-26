package repository

import (
	storageMessage "go-skv/server/storage/message"
)

func switchMessage(terminate func(storageMessage.Terminate), saveRecord func(storageMessage.SaveRecord), forwardToRecord func(msg storageMessage.ForwardToRecord)) func(raw any) (isTerminated bool) {
	return func(raw any) (isTerminated bool) {
		switch msg := raw.(type) {
		case storageMessage.Terminate:
			terminate(msg)
			return true
		case storageMessage.SaveRecord:
			saveRecord(msg)
		case storageMessage.ForwardToRecord:
			forwardToRecord(msg)
		}
		return false
	}
}
