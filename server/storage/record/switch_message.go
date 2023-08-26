package record

import . "go-skv/server/storage/record/message"

func switchMessage(setValue func(SetValue), getValue func(GetValue)) func(any) bool {
	return func(raw any) bool {
		switch msg := raw.(type) {
		case GetValue:
			getValue(msg)
		case SetValue:
			setValue(msg)
		}
		return false
	}
}
