package storagerecord

type RecordDestroyedError struct{}

func (e RecordDestroyedError) Error() string {
	return "record destroyed"
}

type ContextCancelledError struct{}

func (e ContextCancelledError) Error() string {
	return "context cancelled"
}
