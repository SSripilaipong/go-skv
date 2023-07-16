package replicaupdater

import "go-skv/server/dbstorage/dbstoragecontract"

type RecordService interface {
	UpdateReplicaValue(record dbstoragecontract.Record, value string, onFailure func(err error))
	InitializeReplicaRecord(record dbstoragecontract.Record, value string, execute func(record dbstoragecontract.Record))
}
