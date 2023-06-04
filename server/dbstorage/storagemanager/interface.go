package storagemanager

type Interface interface {
	Start() error
	Stop() error
}

type RecordFactory interface {
	New() DbRecord // TODO: manager should destroy all records with ctx when stops
}

type DbRecord interface {
	SetValue(SetValueMessage) error
	GetValue(GetValueMessage) error
	Destroy() error
}
