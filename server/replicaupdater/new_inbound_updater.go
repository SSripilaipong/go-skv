package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/common/util/goutil"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"sync"
)

func (f factory) NewInboundUpdater(ctx context.Context) (replicaupdatercontract.InboundUpdater, error) {
	cmdCh := make(chan actormodel.Command[inboundUpdaterState], 16) // TODO: parameterize buffer size

	initialState := inboundUpdaterState{
		dbStorage: f.dbStorage,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	_ = actormodel.GoRunActor[inboundUpdaterState](ctx, initialState, cmdCh, func(inboundUpdaterState, any) {
		wg.Done()
	})

	return inboundUpdaterInteractor{
		actormodel.NewInteractor[inboundUpdaterState](cmdCh, &wg),
	}, nil
}

type inboundUpdaterState struct {
	dbStorage dbstoragecontract.Storage
}

type inboundUpdaterInteractor struct {
	actormodel.Interactor[inboundUpdaterState]
}

func (t inboundUpdaterInteractor) Update(key string, _ string) error {
	return t.SendCommand(context.Background(), updateInboundReplicaCmd{
		key: key,
	})
}

type updateInboundReplicaCmd struct {
	key string
}

func (c updateInboundReplicaCmd) Execute(state *inboundUpdaterState) {
	goutil.PanicUnhandledError(state.dbStorage.GetRecord(context.Background(), c.key, func(dbstoragecontract.Record) {}))
}
