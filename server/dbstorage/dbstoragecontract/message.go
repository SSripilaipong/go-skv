package dbstoragecontract

type GetRecord struct {
	Key     string
	ReplyTo chan<- any
}
