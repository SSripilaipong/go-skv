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
	goutil.PanicUnhandledError(
		state.dbStorage.GetRecord(context.Background(), c.key, func(record dbstoragecontract.Record) {
			go state.recordService.UpdateReplicaValue(record, c.value, nil)
		}, func(error) {
			go func() {
				goutil.PanicUnhandledError(
					c.sendCommand(context.Background(), createReplicaRecordCmd{key: c.key, value: c.value}),
				)
			}()
		}),
	)
}

type createReplicaRecordCmd struct {
	key   string
	value string
}

func (c createReplicaRecordCmd) Execute(state *inboundUpdaterState) {
	state.recordFactory.New(state.globalCtx)
}
