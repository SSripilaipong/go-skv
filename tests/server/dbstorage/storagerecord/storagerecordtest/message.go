package storagerecordtest

import (
	"go-skv/server/dbstorage/dbstoragecontract"
)

func SendAnyMessage(record dbstoragecontract.Record) error {
	return record.SetValue(nil, "", nil)
}
