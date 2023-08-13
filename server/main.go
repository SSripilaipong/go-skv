package server

import (
	"github.com/urfave/cli/v2"
	"go-skv/common/util/goutil"
	serverCli "go-skv/server/cli"
	"go-skv/server/dbmanager"
	"go-skv/server/dbpeerconnector"
	"go-skv/server/dbserver"
	"go-skv/server/dbstorage"
	"go-skv/server/dbstorage/storagerecord"
	"go-skv/server/replicaupdater"
)

func NewCliCommands() []*cli.Command {
	return serverCli.NewCommands(startServer)
}

func startServer(config serverCli.Config) error {
	storage := dbstorage.New(16, 4)
	replicaUpdaterFactory := replicaupdater.NewFactory(storage, storagerecord.NewFactory(16))
	peerConnector := dbpeerconnector.New(config.PeeringPort, config.AdvertisedIp, config.ExistingPeerAddresses, replicaUpdaterFactory)
	controller := dbserver.New(config.DbPort, storage)

	manager := dbmanager.New(peerConnector, controller, storage)
	goutil.PanicUnhandledError(manager.Start())

	goutil.WaitForInterrupt()

	goutil.PanicUnhandledError(manager.Stop())
	return nil
}
