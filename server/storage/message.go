package storage

type GetValueMessage interface {
	Key() string
	Completed(GetValueResponse) error
}

type SetValueMessage interface {
	Key() string
	Value() string
	Completed(SetValueResponse) error
}
