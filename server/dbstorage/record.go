package dbstorage

type RecordFactory interface {
	New() DbRecord
}

type DbRecord interface {
	SetValue(SetValueMessage) error
	GetValue(GetValueMessage) error
	Destroy() error
}
