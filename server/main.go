package server

import (
	goutil2 "go-skv/common/util/goutil"
	"go-skv/server/dbmanager"
	"go-skv/server/dbpeerconnector"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
	"go-skv/server/servercli"
)

func RunCli() {
	cli := servercli.New(startServer)
	cli.Run()
}

func startServer(config servercli.Config) error {
	storage, storageInteractor := dbstorage.New(16, 4)
	peerConnector := dbpeerconnector.New(config.PeeringPort, config.AdvertisedIp, config.ExistingPeerAddresses)
	controller := dbserver.New(config.DbPort, storageInteractor)

	manager := dbmanager.New(peerConnector, controller, storage)
	goutil2.PanicUnhandledError(manager.Start())

	goutil2.WaitForInterrupt()

	goutil2.PanicUnhandledError(manager.Stop())
	return nil
}
