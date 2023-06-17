package dbpeerconnector

type Interface interface {
	Start() error
	Stop() error
}
