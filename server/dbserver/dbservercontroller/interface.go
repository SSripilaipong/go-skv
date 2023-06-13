package dbservercontroller

type Interface interface {
	Start() error
	Stop() error
	Port() int
}
