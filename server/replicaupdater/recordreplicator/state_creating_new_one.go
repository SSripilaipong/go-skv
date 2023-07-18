package recordreplicator

import "go-skv/common/actormodel"

type creating struct {
	actormodel.Embed
}

func (s *creating) Receive(message any) actormodel.Actor {
	return s
}
