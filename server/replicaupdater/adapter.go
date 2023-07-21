package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/common/commonmessage"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func NewFactoryAdapter(factory replicaupdatercontract.ActorFactory) replicaupdatercontract.Factory {
	return factoryAdapter{factory: factory}
}

type factoryAdapter struct {
	factory replicaupdatercontract.ActorFactory
}

func (a factoryAdapter) NewInboundUpdater(ctx context.Context) (replicaupdatercontract.InboundUpdater, error) {
	ch, err := a.factory.NewInboundUpdater(ctx)
	return adapter{ch: ch}, err
}

type adapter struct {
	ch chan<- any
}

func (a adapter) Update(key string, value string) error {
	go func() {
		a.ch <- InboundUpdate{
			Key:   key,
			Value: value,
		}
	}()
	return nil
}

func (a adapter) Join() {
	panic("deprecated")
}

func newStorageAdapter(dbStorage dbstoragecontract.Storage) chan<- any {
	ch, _ := actormodel.Spawn(context.Background(), &storageAdapter{dbStorage: dbStorage})
	return ch
}

type storageAdapter struct {
	actormodel.Embed
	dbStorage dbstoragecontract.Storage
}

func (s *storageAdapter) Receive(message any) actormodel.Actor {
	switch msg := message.(type) {
	case dbstoragecontract.GetRecord:
		go func() {
			_ = s.dbStorage.GetRecord(context.Background(), msg.Key, func(record dbstoragecontract.Record) {
				defer close(msg.ReplyTo)
				recordCh, _ := actormodel.Spawn(context.Background(), &storagerecord.RecordAdapter{Record: record})
				msg.ReplyTo <- dbstoragecontract.RecordChannel{Ch: recordCh}
			}, func(error) {
				defer close(msg.ReplyTo)
				msg.ReplyTo <- dbstoragecontract.RecordChannel{}
			})
		}()
	case dbstoragecontract.SaveRecord:
		go func() {
			defer close(msg.Ch)
			defer close(msg.ReplyTo)

			ch := make(chan any)
			msg.Ch <- storagerecord.GetRawRecordFromAdapter{ReplyTo: ch}

			record := (<-ch).(dbstoragecontract.Record)
			msg.Ch <- storagerecord.TerminateAdapter{}

			_ = s.dbStorage.Save(context.Background(), msg.Key, record, func(error) {}) // assume success

			msg.ReplyTo <- commonmessage.Ok{Memo: msg.Memo}
		}()
	}
	return s
}
