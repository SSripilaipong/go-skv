package servercli

type dependency struct {
	Start func() error
}
