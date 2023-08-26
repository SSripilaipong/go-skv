package record

import "go-skv/server/storage/record/message"

func switchMessage(getValue func(msg message.GetValue)) func(any) bool {
	return func(raw any) bool {
		switch msg := raw.(type) {
		case message.GetValue:
			getValue(msg)
		}
		return false
	}
}
