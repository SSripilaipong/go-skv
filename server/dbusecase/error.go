package dbusecase

type ContextCancelledError struct{}

func (ContextCancelledError) Error() string {
	return "context cancelled"
}
