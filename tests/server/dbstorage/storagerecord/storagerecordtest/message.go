package storagerecordtest

import (
	"go-skv/server/dbstorage/storagemanager"
	"go-skv/tests/server/dbstorage/storagemanager/storagemanagertest"
)

func SendAnyMessage(record storagemanager.DbRecord) error {
	return record.SetValue(&storagemanagertest.SetValueMessage{})
}
