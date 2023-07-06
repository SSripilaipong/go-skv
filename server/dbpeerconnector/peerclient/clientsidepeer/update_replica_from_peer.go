package clientsidepeer

import (
	"fmt"
	"go-skv/common/util/goutil"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func (t interactor) UpdateReplicaFromPeer(key string, value string) error {
	logUpdateReplicaFromPeer(key, value)
	t.ch <- updateReplicaFromPeerCommand{
		replicaUpdaterFactory: t.replicaUpdaterFactory,
	}
	return nil
}

type updateReplicaFromPeerCommand struct {
	replicaUpdaterFactory replicaupdatercontract.Factory
}

func (c updateReplicaFromPeerCommand) execute(*state) {
	_, err := c.replicaUpdaterFactory.NewInboundUpdater()
	goutil.PanicUnhandledError(err)
}

func logUpdateReplicaFromPeer(key string, value string) {
	fmt.Printf("log: client receive replica (%#v, %#v)\n", key, value)
}
