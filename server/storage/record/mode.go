package record

type recordMode string

const (
	replicaMode recordMode = "REPLICA_MODE"
	ownerMode   recordMode = "OWNER_MODE"
)
