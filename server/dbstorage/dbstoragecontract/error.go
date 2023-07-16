package dbstoragecontract

type RecordDestroyedError struct{}

func (e RecordDestroyedError) Error() string {
	return "record destroyed"
}

type RecordNotFoundError struct{}

func (e RecordNotFoundError) Error() string {
	return "record not found"
}

type DuplicateKeyError struct{}

func (e DuplicateKeyError) Error() string {
	return "duplicate key"
}
