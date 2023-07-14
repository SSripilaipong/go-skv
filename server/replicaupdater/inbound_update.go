package replicaupdater

import (
	"context"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func (t inboundUpdaterInteractor) Update(key string, value string) error {
	return t.SendCommand(context.Background(), updateInboundReplicaCmd{
		key:   key,
		value: value,
	})
}

type updateInboundReplicaCmd struct {
	key   string
	value string
}

func (c updateInboundReplicaCmd) Execute(state *inboundUpdaterState) {
	goutil.PanicUnhandledError(
		state.dbStorage.GetRecord(context.Background(), c.key, func(record dbstoragecontract.Record) {
			go state.recordService.UpdateReplicaValue(record, c.value, nil)
		}),
	)
}
