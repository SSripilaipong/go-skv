package message

type GetValue struct {
	ReplyTo chan any
	Memo    string
}

type SetValue struct {
	Value   string
	Memo    string
	ReplyTo chan any
}

type Ack struct {
	Memo string
}

type Value struct {
	Value string
	Memo  string
}
