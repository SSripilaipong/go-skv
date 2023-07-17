package actormodel

import (
	"sync"
)

type ActorRef interface {
	Join()
	channel() chan packet
}

type actorRef struct {
	ch chan packet
	wg *sync.WaitGroup
}

func (r actorRef) channel() chan packet {
	return r.ch
}

func (r actorRef) Join() {
	r.wg.Wait()
}
