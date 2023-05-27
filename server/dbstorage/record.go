package dbstorage

type DbRecord interface {
	SetValue(value string) error
}
