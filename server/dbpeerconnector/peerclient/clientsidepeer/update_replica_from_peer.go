package clientsidepeer

import (
	"fmt"
	"go-skv/common/util/goutil"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func (t interactor) UpdateReplicaFromPeer(key string, value string) error {
	logUpdateReplicaFromPeer(key, value)
	t.ch <- updateReplicaFromPeerCommand{
		key:                   key,
		value:                 value,
		replicaUpdaterFactory: t.replicaUpdaterFactory,
	}
	return nil
}

type updateReplicaFromPeerCommand struct {
	key                   string
	value                 string
	replicaUpdaterFactory replicaupdatercontract.Factory
}

func (c updateReplicaFromPeerCommand) execute(s *state) {
	updater := s.inboundUpdater
	if updater == nil {
		var err error
		updater, err = c.replicaUpdaterFactory.NewInboundUpdater(s.ctx)
		goutil.PanicUnhandledError(err)
		s.inboundUpdater = updater
	}

	goutil.PanicUnhandledError(updater.Update(c.key, c.value))
}

func logUpdateReplicaFromPeer(key string, value string) {
	fmt.Printf("log: client receive replica (%#v, %#v)\n", key, value)
}
