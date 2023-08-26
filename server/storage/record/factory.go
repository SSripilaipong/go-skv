package record

import (
	"context"
	"go-skv/common/util/goutil"
)

type factory struct {
	channelBufferSize int
}

func (f factory) New(ctx context.Context, value string) chan<- any {
	ch := make(chan any, f.channelBufferSize)

	go loop(ctx, switchMessage(
		getValue(ctx, &value),
	), ch)

	return goutil.ExtendedSenderChannel(ch)
}
