package storagerecordtest

import (
	"go-skv/server/dbstorage/storagerecord"
)

func SendAnyMessage(record storagerecord.Interface) error {
	return record.SetValue("", nil)
}
