package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (r recordInteractor) SetValue(ctx context.Context, value string, success func(response dbstoragecontract.RecordData)) error {
	return r.sendCommand(ctx, setValueCommand{
		value:   value,
		success: success,
	})
}

type setValueCommand struct {
	value   string
	success func(dbstoragecontract.RecordData)
}

func (c setValueCommand) execute(s *state) {
	s.value = c.value
	c.success(dbstoragecontract.RecordData{Value: c.value})
}
