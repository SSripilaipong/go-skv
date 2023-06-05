package storagerecord

type RecordDestroyedError struct{}

func (e RecordDestroyedError) Error() string {
	return "record destroyed"
}
