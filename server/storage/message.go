package storage

type GetValueMessage struct {
	Key string
}

type Packet interface {
	Message() any
}
