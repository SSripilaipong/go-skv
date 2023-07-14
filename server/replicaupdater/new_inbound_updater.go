package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
	"go-skv/server/dbstorage/dbstoragecontract"
	"go-skv/server/replicaupdater/replicaupdatercontract"
	"sync"
)

func (f factory) NewInboundUpdater(ctx context.Context) (replicaupdatercontract.InboundUpdater, error) {
	cmdCh := make(chan actormodel.Command[inboundUpdaterState], 16) // TODO: parameterize buffer size

	initialState := inboundUpdaterState{
		globalCtx:     ctx,
		dbStorage:     f.dbStorage,
		recordService: f.recordService,
		recordFactory: f.recordFactory,
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
	globalCtx     context.Context
	dbStorage     dbstoragecontract.Storage
	recordService RecordService
	recordFactory dbstoragecontract.Factory
}

type inboundUpdaterInteractor struct {
	actormodel.Interactor[inboundUpdaterState]
}

var _ replicaupdatercontract.InboundUpdater = &inboundUpdaterInteractor{}
