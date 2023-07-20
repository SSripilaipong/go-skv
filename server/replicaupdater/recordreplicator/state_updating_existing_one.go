package recordreplicator

import (
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type updating struct {
	actormodel.Embed
	recordFactory dbstoragecontract.Factory
	storage       chan<- any
	key           string
	value         string
}

func (s *updating) Receive(message any) actormodel.Actor {
	switch msg := message.(type) {
	case dbstoragecontract.RecordChannel:
		return s.receiveResponseFromRepository(msg)
	case commonmessage.Ok:
		return nil
	default:
		return s
	}
}

func (s *updating) receiveResponseFromRepository(msg dbstoragecontract.RecordChannel) actormodel.Actor {
	if msg.Ch != nil {
		defer close(msg.Ch)

		if sent := s.SendIfNotDone(msg.Ch, dbstoragecontract.UpdateReplicaValue{
			Value:   s.value,
			ReplyTo: s.Self(),
		}); !sent {
			return nil
		}
		return s
	}

	s.ScheduleReceive(commonmessage.Start{})
	return &creating{
		recordFactory: s.recordFactory,
		storage:       s.storage,
		key:           s.key,
	}
}
