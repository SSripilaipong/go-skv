package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
)

func SendAnyMessage(record storagerecord.DbRecord) error {
	return record.SetValue(&storagemanagertest.SetValueMessage{})
}
