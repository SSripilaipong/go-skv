package actormodel

import (
	"context"
	"reflect"
)

func NewTestActor(options ...func(*spawnParams)) *TestActorRef {
	params := spawnParams{
		bufferSize: 8,
	}

	for _, option := range options {
		option(&params)
	}
	ch := make(chan packet, params.bufferSize)
	return &TestActorRef{actorRef{ch: ch}}
}

type TestActorRef struct {
	actorRef
}

func (r *TestActorRef) FakeTellBlocking(ctx context.Context, receiver ActorRef, message any) error {
	return tellBlocking(ctx, receiver.channel(), r.channel(), message)
}

func (r *TestActorRef) SeekMessage(ctx context.Context, t any) (any, bool) {
	for {
		select {
		case pk := <-r.ch:
			if reflect.TypeOf(pk.message) == reflect.TypeOf(t) {
				return pk.message, true
			}
		case <-ctx.Done():
			return nil, false
		}
	}
}
