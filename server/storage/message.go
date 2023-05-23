package storage

type GetValueMessage interface {
	Key() string
	Completed(GetValueResponse) error
}
