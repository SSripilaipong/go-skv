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

	goRecordActor(ctx, ownerMode, value, ch)

	return goutil.ExtendedSenderChannel(ch)
}

func (f factory) NewReplica(ctx context.Context, value string) chan<- any {
	ch := make(chan any, f.channelBufferSize)

	goRecordActor(ctx, replicaMode, value, ch)

	return goutil.ExtendedSenderChannel(ch)
}

func goRecordActor(ctx context.Context, mode recordMode, value string, ch chan any) {
	go loop(ctx, switchMode(mode, switchOwnerMessage(
		setValue(ctx, &value),
		getValue(ctx, &value),
	), switchReplicaMessage(
		getValue(ctx, &value),
	)), ch)
}
