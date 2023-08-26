package record

import . "go-skv/server/storage/record/message"

func switchOwnerMessage(setValue func(SetValue), getValue func(GetValue)) func(any) (recordMode, bool) {
	return func(raw any) (recordMode, bool) {
		switch msg := raw.(type) {
		case GetValue:
			getValue(msg)
		case SetValue:
			setValue(msg)
		}
		return ownerMode, false
	}
}
