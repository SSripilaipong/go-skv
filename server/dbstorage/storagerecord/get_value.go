package storagerecord

import (
	"context"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (r recordInteractor) GetValue(ctx context.Context, success func(response dbstoragecontract.RecordData)) error {
	return r.sendCommand(ctx, getValueCommand{
		success: success,
	})
}

type getValueCommand struct {
	success func(dbstoragecontract.RecordData)
}

func (c getValueCommand) execute(s *state) {
	c.success(dbstoragecontract.RecordData{Value: s.value})
}
