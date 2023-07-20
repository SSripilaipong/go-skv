package recordreplicator

import (
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

const setRecordMemo = "set record mode"

type creating struct {
	actormodel.Embed
	recordFactory dbstoragecontract.Factory
	storage       chan<- any
	key           string
}

func (s *creating) Receive(message any) actormodel.Actor {
	switch message.(type) {
	case commonmessage.Start:
		createdRecord := s.recordFactory.NewActor(s.Ctx())
		if sent := s.SendIfNotDone(createdRecord, dbstoragecontract.SetRecordMode{
			Mode:    dbstoragecontract.ReplicaMode,
			Memo:    setRecordMemo,
			ReplyTo: s.Self(),
		}); !sent {
			return nil
		}

		return &creating_settingRecordMode{
			storage:       s.storage,
			key:           s.key,
			createdRecord: createdRecord,
		}
	}
	return s
}

type creating_settingRecordMode struct {
	actormodel.Embed
	storage       chan<- any
	key           string
	createdRecord chan<- any
}

func (s *creating_settingRecordMode) Receive(message any) actormodel.Actor {
	switch msg := message.(type) {
	case commonmessage.Ok:
		if msg.Memo == setRecordMemo {
			if sent := s.SendIfNotDone(s.storage, dbstoragecontract.SaveRecord{
				Key: s.key,
				Ch:  s.createdRecord,
			}); !sent {
				return nil
			}
		}
	}
	return s
}
