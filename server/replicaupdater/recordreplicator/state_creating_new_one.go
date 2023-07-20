package recordreplicator

import (
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type creating struct {
	actormodel.Embed
	recordFactory dbstoragecontract.Factory
}

func (s *creating) Receive(message any) actormodel.Actor {
	switch message.(type) {
	case commonmessage.Start:
		createdRecord := s.recordFactory.NewActor(s.Ctx())
		if sent := s.SendIfNotDone(createdRecord, dbstoragecontract.SetRecordMode{Mode: dbstoragecontract.ReplicaMode}); !sent {
			return nil
		}
	}
	return s
}
