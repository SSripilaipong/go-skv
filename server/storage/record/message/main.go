package message

type GetValue struct {
	ReplyTo chan any
}

type Value struct {
	Value string
}
