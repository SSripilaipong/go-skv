package commoncontract

type ContextClosedError struct{}

func (e ContextClosedError) Error() string {
	return "context closed"
}
