package repositoryroutine

type Interface interface {
	Start() error
	Stop() error
}
