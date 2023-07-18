package recordreplicator

import (
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type updatingState struct {
	actormodel.Embed
	value string
}

func (s *updatingState) Receive(message any) actormodel.Actor {
	switch msg := message.(type) {
	case dbstoragecontract.RecordChannel:
		msg.Ch <- dbstoragecontract.UpdateReplicaValue{Value: s.value, ReplyTo: s.Self()}
		return s
	case commonmessage.Ok:
		return nil
	default:
		return s
	}
}
