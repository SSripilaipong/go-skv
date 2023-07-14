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
		dbStorage:     f.dbStorage,
		recordService: f.recordService,
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
	dbStorage     dbstoragecontract.Storage
	recordService RecordService
}

type inboundUpdaterInteractor struct {
	actormodel.Interactor[inboundUpdaterState]
}

var _ replicaupdatercontract.InboundUpdater = &inboundUpdaterInteractor{}
