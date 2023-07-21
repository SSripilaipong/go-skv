package storagerecord

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
)

type GetRawRecordFromAdapter struct {
	ReplyTo chan<- any
}

type TerminateAdapter struct{}

type RecordAdapter struct {
	actormodel.Embed
	Record dbstoragecontract.Record
}

func (s *RecordAdapter) Receive(message any) actormodel.Actor {
	switch msg := message.(type) {
	case dbstoragecontract.UpdateReplicaValue:
		go func() {
			_ = s.Record.SetValue(context.Background(), msg.Value, func(dbstoragecontract.RecordData) {
				defer close(msg.ReplyTo)
				msg.ReplyTo <- commonmessage.Ok{Memo: msg.Memo}
			})
		}()
	case dbstoragecontract.SetRecordMode:
		go func() {
			defer close(msg.ReplyTo)
			msg.ReplyTo <- commonmessage.Ok{Memo: msg.Memo}
		}()
	case GetRawRecordFromAdapter:
		go func() {
			defer close(msg.ReplyTo)
			msg.ReplyTo <- s.Record
		}()
	case TerminateAdapter:
		return nil
	}
	return s
}
