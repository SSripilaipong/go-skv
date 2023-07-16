package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (t inboundUpdaterInteractor) Update(key string, value string) error {
	return t.SendCommand(context.Background(), actormodel.Do(t.updateInboundReplicaCmd(key, value)))
}

func (t inboundUpdaterInteractor) updateInboundReplicaCmd(key, value string) func(state *inboundUpdaterState) {
	return func(state *inboundUpdaterState) {

		tryUpdateReplicaRecord := func(record dbstoragecontract.Record) {
			go state.recordService.UpdateReplicaValue(record, value, nil)
		}

		createNewReplicaRecordIfNotExists := func(error) {
			go t.SendCommandOrPanic(actormodel.Do(t.createReplicaRecordCmd(key, value)))
		}

		goutil.PanicUnhandledError(
			state.dbStorage.GetRecord(context.Background(), key, tryUpdateReplicaRecord, createNewReplicaRecordIfNotExists),
		)
	}
}

func (t inboundUpdaterInteractor) createReplicaRecordCmd(key, value string) func(state *inboundUpdaterState) {
	return func(state *inboundUpdaterState) {
		record := state.recordFactory.New(state.globalCtx)
		state.recordService.InitializeReplicaRecord(record, value, func(record dbstoragecontract.Record) {
			goutil.PanicUnhandledError(
				state.dbStorage.Save(context.Background(), key, record),
			)
		})
	}
}
