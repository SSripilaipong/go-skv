package repository

import (
	"context"
	"go-skv/common/util/goutil"
)

func newRepository(ctx context.Context, bufferSize int16) chan<- any {
	ch := make(chan any, bufferSize)
	go loop(ctx, switchMessage(
		terminate(ctx),
	), ch)
	return goutil.ExtendedSenderChannel(ch)
}
