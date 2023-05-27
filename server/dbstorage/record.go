package dbstorage

type DbRecord interface {
	SetValue(SetValueMessage) error
	GetValue(message GetValueMessage) error
}
