package dbstorage

type DbRecord interface {
	SetValue(SetValueMessage) error
}
