package repositoryinteractor

type ContextCancelledError struct {
}

func (e ContextCancelledError) Error() string {
	return "context cancelled error"
}
