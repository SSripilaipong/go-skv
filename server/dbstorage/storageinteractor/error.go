package storageinteractor

type TimeoutError struct {
}

func (e TimeoutError) Error() string {
	return "timeout error"
}
