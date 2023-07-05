package server

import (
	"go-skv/common/util/goutil"
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
	storage := dbstorage.New(16, 4)
	peerConnector := dbpeerconnector.New(config.PeeringPort, config.AdvertisedIp, config.ExistingPeerAddresses)
	controller := dbserver.New(config.DbPort, storage)

	manager := dbmanager.New(peerConnector, controller, storage)
	goutil.PanicUnhandledError(manager.Start())

	goutil.WaitForInterrupt()

	goutil.PanicUnhandledError(manager.Stop())
	return nil
}
