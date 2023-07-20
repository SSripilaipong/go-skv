package dbstoragecontract

type GetRecord struct {
	Key     string
	ReplyTo chan<- any
}

type RecordChannel struct {
	Ch chan<- any
}

// record messages -> should move to a separate file

type UpdateReplicaValue struct {
	Value   string
	ReplyTo chan<- any
}

type RecordMode uint8

const (
	OwnerMode RecordMode = iota
	ReplicaMode
)

type SetRecordMode struct {
	Mode RecordMode
}
