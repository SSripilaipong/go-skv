package storagerecord

import (
	"context"
)

func (r recordInteractor) SetValue(ctx context.Context, value string, success func(response SetValueResponse)) error {
	return r.sendCommand(ctx, setValueCommand{
		value:   value,
		success: success,
	})
}

type SetValueResponse struct {
	Value string
}

type setValueCommand struct {
	value   string
	success func(SetValueResponse)
}

func (c setValueCommand) execute(s *state) {
	s.value = c.value
	c.success(SetValueResponse{Value: c.value})
}
