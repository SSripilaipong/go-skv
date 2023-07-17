package actormodel

import (
	"context"
	"go-skv/common/commoncontract"
	"sync"
)

type Actor interface {
	Receive(sender ActorRef, message any) Actor
	setProps(ch chan packet, ref ActorRef)
}

type Embed struct {
	ch  chan packet
	ref ActorRef
}

func (t *Embed) setProps(ch chan packet, ref ActorRef) {
	t.ch = ch
	t.ref = ref
}

func (t *Embed) Ref() ActorRef {
	return t.ref
}

func (t *Embed) TellBlocking(ctx context.Context, receiver ActorRef, message any) error {
	return tellBlocking(ctx, receiver.channel(), t.ch, message)
}

type packet struct {
	sender  ActorRef
	message any
}

type spawnParams struct {
	bufferSize int
}

func Spawn(ctx context.Context, actor Actor, options ...func(*spawnParams)) ActorRef {
	params := spawnParams{
		bufferSize: 0,
	}

	for _, option := range options {
		option(&params)
	}

	ch := make(chan packet, params.bufferSize)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ref := actorRef{ch: ch, wg: wg}
	go runActorLoop(ctx, ch, wg, actor, ref)
	return ref
}

func runActorLoop(ctx context.Context, ch chan packet, wg *sync.WaitGroup, actor Actor, ref actorRef) {
	defer func() { wg.Done() }()

	for {
		select {
		case pk := <-ch:
			actor.setProps(ch, ref)
			actor = actor.Receive(pk.sender, pk.message)
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

func tellBlocking(ctx context.Context, recvCh chan packet, sendCh chan packet, message any) error {
	select {
	case recvCh <- packet{sender: actorRef{ch: sendCh}, message: message}:
		return nil
	case <-ctx.Done():
		return commoncontract.ContextClosedError{}
	}
}
