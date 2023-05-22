package dbmanager

type DbServer interface {
	Start() error
	Stop() error
	Port() int
}
