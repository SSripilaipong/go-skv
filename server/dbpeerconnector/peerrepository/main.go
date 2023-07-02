package peerrepository

import (
	"context"
	"errors"
	"go-skv/server/dbpeerconnector/peerconnectorcontract"
	"go-skv/server/dbpeerconnector/peerrepository/peerrepositorycontract"
	"go-skv/util/goutil"
)

func New(ctx context.Context) peerrepositorycontract.Repository {
	ch := make(chan command)
	go mainLoop(ctx, ch)
	return interactor{ch: ch}
}

type interactor struct {
	ch chan<- command
}

func (t interactor) sendCommand(ctx context.Context, cmd command) {
	select {
	case t.ch <- cmd:
	case <-ctx.Done():
		goutil.PanicUnhandledError(errors.New("context closed"))
	}
}

type state struct {
	temp peerconnectorcontract.Peer
}

type command interface {
	execute(s *state)
}

func mainLoop(ctx context.Context, ch <-chan command) {
	var s state
	var cmd command
	for {
		select {
		case cmd = <-ch:
			cmd.execute(&s)
		case <-ctx.Done():
			return
		}
	}
}
