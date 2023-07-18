package actormodel

import (
	"context"
	"go-skv/common/commoncontract"
	"sync"
)

type Actor interface {
	Receive(message any) Actor
	setProps(ctx context.Context, ch chan any)
}

type Embed struct {
	ctx context.Context
	ch  chan any
}

func (t *Embed) setProps(ctx context.Context, ch chan any) {
	t.ctx = ctx
	t.ch = ch
}

func (t *Embed) Self() chan<- any {
	return t.ch
}

func (t *Embed) Ctx() context.Context {
	return t.ctx
}

func (t *Embed) TellBlocking(ctx context.Context, receiver chan<- any, message any) error {
	return tellBlocking(ctx, receiver, message)
}

type spawnParams struct {
	bufferSize int
}

func Spawn(ctx context.Context, actor Actor, options ...func(*spawnParams)) (chan<- any, func()) {
	params := spawnParams{
		bufferSize: 0,
	}

	for _, option := range options {
		option(&params)
	}

	ch := make(chan any, params.bufferSize)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go runActorLoop(ctx, ch, wg, actor)
	return ch, wg.Wait
}

func runActorLoop(ctx context.Context, ch chan any, wg *sync.WaitGroup, actor Actor) {
	defer func() { wg.Done() }()

	for {
		select {
		case message := <-ch:
			actor.setProps(ctx, ch)
			if actor = actor.Receive(message); actor == nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func WithBufferSize(size int) func(*spawnParams) {
	return func(params *spawnParams) {
		params.bufferSize = size
	}
}

func tellBlocking(ctx context.Context, recvCh chan<- any, message any) error {
	select {
	case recvCh <- message:
		return nil
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
}

type assertTypeEmbedActor struct{ Embed }

func (assertTypeEmbedActor) Receive(any) Actor { return nil }

var _ Actor = &assertTypeEmbedActor{}
