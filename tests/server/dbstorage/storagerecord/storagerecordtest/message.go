package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/storagerepository/repositoryroutine/repositoryroutinetest"
)

func SendAnyMessage(record storagerecord.DbRecord) error {
	return record.SetValue(&repositoryroutinetest.SetValueMessage{})
}
