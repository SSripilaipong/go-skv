package storagerepository

type Interface interface {
	Start() error
	Stop() error
}
