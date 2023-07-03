package storagerecord

import (
	"context"
)

func (r recordInteractor) GetValue(ctx context.Context, success func(response GetValueResponse)) error {
	return r.sendCommand(ctx, getValueCommand{
		success: success,
	})
}

type GetValueResponse struct {
	Value string
}

type getValueCommand struct {
	success func(GetValueResponse)
}

func (c getValueCommand) execute(s *state) {
	c.success(GetValueResponse{Value: s.value})
}
