package servercli

type dependency struct {
	Start func(config Config) error
}
