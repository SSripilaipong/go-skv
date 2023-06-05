package repositoryinteractor

type TimeoutError struct {
}

func (e TimeoutError) Error() string {
	return "timeout error"
}
