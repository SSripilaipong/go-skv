package record

func switchMode(initialMode recordMode, ownerHandle, replicaHandle func(any) (recordMode, bool)) func(any) bool {
	mode := initialMode
	var isTerminated bool
	return func(raw any) bool {
		switch mode {
		case ownerMode:
			mode, isTerminated = ownerHandle(raw)
		case replicaMode:
			mode, isTerminated = replicaHandle(raw)
		}
		return isTerminated
	}
}
