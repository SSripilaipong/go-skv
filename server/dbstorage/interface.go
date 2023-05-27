package dbstorage

type Interface interface {
	Start() error
	Stop() error
}
