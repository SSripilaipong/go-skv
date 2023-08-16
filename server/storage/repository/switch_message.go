package repository

import (
	storageMessage "go-skv/server/storage/message"
)

func switchMessage(terminate func(terminate storageMessage.Terminate)) func(raw any) (isTerminated bool) {
	return func(raw any) (isTerminated bool) {
		switch msg := raw.(type) {
		case storageMessage.Terminate:
			terminate(msg)
			return true
		}
		return false
	}
}
