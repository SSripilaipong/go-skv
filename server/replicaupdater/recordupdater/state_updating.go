package recordupdater

import "go-skv/common/actormodel"

type updatingState struct {
	actormodel.Embed
}

func (s *updatingState) Receive(message any) actormodel.Actor {
	return nil
}
