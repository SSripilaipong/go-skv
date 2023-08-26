package repository

import (
	"context"
	"go-skv/common/util/goutil"
)

func newRepository(ctx context.Context, bufferSize int16) chan<- any {
	ch := make(chan any, bufferSize)

	records := make(map[string]chan<- any)

	go loop(ctx, switchMessage(
		terminate(ctx),
		saveRecord(ctx, records),
		forwardToRecord(ctx, records),
	), ch)
	return goutil.ExtendedSenderChannel(ch)
}
