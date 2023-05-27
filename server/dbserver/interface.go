package dbserver

type Interface interface {
	Start() error
	Stop() error
	Port() int
}
