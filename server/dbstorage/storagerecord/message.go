package storagerecord

type getValueMessage struct {
	success func(GetValueResponse)
}

type setValueMessage struct {
	value   string
	success func(SetValueResponse)
}
