package message

type Terminate struct {
	Notify chan<- struct{}
}

type Ack struct {
	Memo string
}

type SaveRecord struct {
	Key     string
	Channel chan<- any
	Memo    string
	ReplyTo chan<- any
}

type ForwardToRecord struct {
	Key     string
	Message any
	Memo    string
	ReplyTo chan<- any
}
