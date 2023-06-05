package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagerepository/storagerepositorytest"
)

func SendAnyMessage(record storagerecord.DbRecord) error {
	return record.SetValue(&storagerepositorytest.SetValueMessage{})
}
