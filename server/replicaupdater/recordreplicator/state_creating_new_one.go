package recordreplicator

import (
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type creating struct {
	actormodel.Embed
	recordFactory dbstoragecontract.Factory
	storage       chan<- any
	key           string
	value         string
}

func (s *creating) Receive(message any) actormodel.Actor {
	switch message.(type) {
	case commonmessage.Start:
		createdRecord := s.recordFactory.NewActor(s.Ctx())
		if sent := s.SendIfNotDone(createdRecord, dbstoragecontract.SetRecordMode{
			Mode:    dbstoragecontract.ReplicaMode,
			Memo:    setRecordModeMemo,
			ReplyTo: s.Self(),
		}); !sent {
			return nil
		}

		return &creating_settingRecordMode{
			storage:       s.storage,
			key:           s.key,
			value:         s.value,
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
	value         string
}

func (s *creating_settingRecordMode) Receive(message any) actormodel.Actor {
	switch msg := message.(type) {
	case commonmessage.Ok:
		if msg.Memo == setRecordModeMemo {
			if sent := s.SendIfNotDone(s.storage, dbstoragecontract.SaveRecord{
				Key:  s.key,
				Ch:   s.createdRecord,
				Memo: saveRecordMemo,
			}); !sent {
				return nil
			}
			return &creating_savingRecord{
				createdRecord: s.createdRecord,
				value:         s.value,
			}
		}
	}
	return s
}

type creating_savingRecord struct {
	actormodel.Embed
	value         string
	createdRecord chan<- any
}

func (s *creating_savingRecord) Receive(message any) actormodel.Actor {
	switch message.(type) {
	case commonmessage.Ok:
		s.ScheduleReceive(dbstoragecontract.RecordChannel{Ch: s.createdRecord})
		return &updating{value: s.value}
	}
	return s
}
