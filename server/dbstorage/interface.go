package dbstorage

type Interface interface {
	Start() error
	Stop() error
}

type RecordFactory interface {
	New() DbRecord
}
