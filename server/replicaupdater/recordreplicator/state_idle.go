package recordreplicator

import (
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type idle struct {
	actormodel.Embed
	storage       chan<- any
	recordFactory dbstoragecontract.Factory
	key           string
	value         string
}

func (s *idle) Receive(message any) actormodel.Actor {
	if _, isStartMessage := message.(commonmessage.Start); isStartMessage {
		if sent := s.SendIfNotDone(s.storage, dbstoragecontract.GetRecord{
			Key:     s.key,
			ReplyTo: s.Self(),
		}); !sent {
			return nil
		}

		return &updating{
			recordFactory: s.recordFactory,
			value:         s.value,
		}
	}
	return s
}
