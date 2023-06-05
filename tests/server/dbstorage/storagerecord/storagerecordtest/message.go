package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/tests/server/dbstorage/repositoryroutine/repositoryroutinetest"
)

func SendAnyMessage(record storagerecord.Interface) error {
	return record.SetValue(&repositoryroutinetest.SetValueMessage{})
}
