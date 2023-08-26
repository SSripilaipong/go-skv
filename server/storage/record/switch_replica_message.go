package record

import . "go-skv/server/storage/record/message"

func switchReplicaMessage(getValue func(GetValue)) func(any) (recordMode, bool) {
	return func(raw any) (recordMode, bool) {
		switch msg := raw.(type) {
		case GetValue:
			getValue(msg)
		}
		return replicaMode, false
	}
}
