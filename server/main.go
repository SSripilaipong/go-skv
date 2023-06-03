package server

import (
	"context"
	"go-skv/server/dbmanager"
	"go-skv/server/dbpeerserver"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/dbstoragerecord"
	"go-skv/server/dbusecase"
	"go-skv/server/servercli"
	"go-skv/util/goutil"
	"os"
	"os/signal"
)

func RunCli() {
	cli := servercli.New(startServer)
	cli.Run()
}

func startServer() error {
	ctx, cancelCtx := context.WithCancel(context.Background())
	storageChan := make(chan any, 16)
	storage := dbstorage.New(ctx, storageChan, dbstoragerecord.NewFactory(ctx, 4))
	usecaseDep := dbusecase.NewDependency(storageChan)
	manager := dbmanager.New(dbpeerserver.New(), dbserver.New(5555, dbserver.Dependency{
		GetValueUsecase: dbusecase.GetValueUsecase(usecaseDep),
		SetValueUsecase: dbusecase.SetValueUsecase(usecaseDep),
	}), storage)
	goutil.PanicUnhandledError(manager.Start())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	goutil.PanicUnhandledError(manager.Stop())
	cancelCtx()
	return nil
}
