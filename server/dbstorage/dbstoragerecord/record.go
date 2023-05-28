package dbstoragerecord

import (
	"go-skv/goutil"
	"go-skv/server/dbstorage"
)

type recordInterface struct {
}

func newRecordInterface() dbstorage.DbRecord {
	return &recordInterface{}
}

func (r *recordInterface) SetValue(message dbstorage.SetValueMessage) error {
	goutil.PanicUnhandledError(message.Completed(dbstorage.SetValueResponse{}))
	return nil
}

func (r *recordInterface) GetValue(dbstorage.GetValueMessage) error {
	return nil
}

func (r *recordInterface) Destroy() error {
	return nil
}
