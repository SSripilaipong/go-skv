package clientrepl

type ReplClosedError struct{}

func (e ReplClosedError) Error() string {
	return "repl closed"
}
