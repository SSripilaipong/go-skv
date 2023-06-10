package storagerecord

type GetValueMessage interface {
	Key() string
	Completed(GetValueResponse) error
}

type SetValueMessage interface {
	Key() string
	Value() string
	Completed(SetValueResponse) error
}

type setValueMessage struct {
	key     string
	value   string
	success func(SetValueResponse)
}
