package recordreplicator

import (
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type idleState struct {
	actormodel.Embed
	storage chan<- any
	key     string
	value   string
}

func (s *idleState) Receive(message any) actormodel.Actor {
	if _, isStartMessage := message.(commonmessage.Start); isStartMessage {
		s.storage <- dbstoragecontract.GetRecord{Key: s.key, ReplyTo: s.Self()}
		return &updatingState{value: s.value}
	}
	return s
}
