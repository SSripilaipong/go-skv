package clientsidepeer

import (
	"go-skv/common/util/goutil"
	"go-skv/server/replicaupdater/replicaupdatercontract"
)

func (t interactor) UpdateReplicaFromPeer(key string, value string) error {
	return t.sendCommand(updateReplicaFromPeerCommand{
		key:                   key,
		value:                 value,
		replicaUpdaterFactory: t.replicaUpdaterFactory,
	})
}

type updateReplicaFromPeerCommand struct {
	key                   string
	value                 string
	replicaUpdaterFactory replicaupdatercontract.ActorFactory
}

func (c updateReplicaFromPeerCommand) execute(s *state) {
	updater := s.inboundUpdater
	if updater == nil {
		var err error
		updater, err = c.replicaUpdaterFactory.NewInboundUpdater(s.ctx)
		goutil.PanicUnhandledError(err)
		s.inboundUpdater = updater
	}

	select {
	case updater <- replicaupdatercontract.InboundUpdate{Key: c.key, Value: c.value}:
	case <-s.ctx.Done():
	}
}
