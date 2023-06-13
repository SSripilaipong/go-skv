package server

import (
	"go-skv/server/dbmanager"
	"go-skv/server/dbpeerconnector"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
	"go-skv/server/servercli"
	"go-skv/util/goutil"
)

func RunCli() {
	cli := servercli.New(startServer)
	cli.Run()
}

func startServer() error {
	storage, storageInteractor := dbstorage.New(16, 4)
	usecaseDep := dbusecase.NewDependency(storageInteractor)
	peerConnector := dbpeerconnector.New()
	rpcServer := dbserver.New(5555, dbserver.Dependency{
		GetValueUsecase: dbusecase.GetValueUsecase(usecaseDep),
		SetValueUsecase: dbusecase.SetValueUsecase(usecaseDep),
	})

	manager := dbmanager.New(peerConnector, rpcServer, storage)
	goutil.PanicUnhandledError(manager.Start())

	goutil.WaitForInterrupt()

	goutil.PanicUnhandledError(manager.Stop())
	return nil
}
