package storagemanager

type Interface interface {
	Start() error
	Stop() error
}
