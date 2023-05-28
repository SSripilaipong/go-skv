package dbstoragerecordtest

import (
	"go-skv/server/dbstorage"
	"go-skv/tests/server/dbstorage/dbstoragetest"
)

func SendAnyMessage(record dbstorage.DbRecord) error {
	return record.SetValue(&dbstoragetest.SetValueMessage{})
}
