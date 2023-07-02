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
	initialState := state{
		peers: make(map[string]peerconnectorcontract.Peer),
	}
	go mainLoop(ctx, initialState, ch)
	return interactor{ch: ch}
}

type state struct {
	peers map[string]peerconnectorcontract.Peer
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

type command interface {
	execute(s *state)
}

func mainLoop(ctx context.Context, s state, ch <-chan command) {
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
