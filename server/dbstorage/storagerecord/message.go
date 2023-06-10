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

type getValueMessage struct {
	success func(GetValueResponse)
}

type setValueMessage struct {
	value   string
	success func(SetValueResponse)
}
