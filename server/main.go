package server

import (
	"go-skv/server/dbmanager"
	"go-skv/server/dbpeerserver"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
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
	storage, storageChan := dbstorage.New(16, 4)
	usecaseDep := dbusecase.NewDependency(storageChan)
	peerServer := dbpeerserver.New()
	rpcServer := dbserver.New(5555, dbserver.Dependency{
		GetValueUsecase: dbusecase.GetValueUsecase(usecaseDep),
		SetValueUsecase: dbusecase.SetValueUsecase(usecaseDep),
	})

	manager := dbmanager.New(peerServer, rpcServer, storage)
	goutil.PanicUnhandledError(manager.Start())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	goutil.PanicUnhandledError(manager.Stop())
	return nil
}
