package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (t inboundUpdaterInteractor) Update(key string, value string) error {
	return t.SendCommand(context.Background(), updateInboundReplicaCmd{
		key:         key,
		value:       value,
		sendCommand: t.SendCommand,
	})
}

type updateInboundReplicaCmd struct {
	key         string
	value       string
	sendCommand func(ctx context.Context, cmd actormodel.Command[inboundUpdaterState]) error
}

func (c updateInboundReplicaCmd) Execute(state *inboundUpdaterState) {
	sendCommandToSelf := func(cmd actormodel.Command[inboundUpdaterState]) {
		goutil.PanicUnhandledError(c.sendCommand(context.Background(), cmd))
	}

	tryUpdateReplicaRecord := func(record dbstoragecontract.Record) {
		go state.recordService.UpdateReplicaValue(record, c.value, nil)
	}

	createNewReplicaRecordIfNotExists := func(error) {
		go sendCommandToSelf(createReplicaRecordCmd{value: c.value})
	}

	goutil.PanicUnhandledError(
		state.dbStorage.GetRecord(context.Background(), c.key, tryUpdateReplicaRecord, createNewReplicaRecordIfNotExists),
	)
}

type createReplicaRecordCmd struct {
	value string
}

func (c createReplicaRecordCmd) Execute(state *inboundUpdaterState) {
	record := state.recordFactory.New(state.globalCtx)
	state.recordService.InitializeReplicaRecord(record, c.value, nil)
}
