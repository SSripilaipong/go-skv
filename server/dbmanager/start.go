package dbmanager

import (
	"go-skv/common/util/goutil"
)

func (m manager) Start() error {
	goutil.PanicUnhandledError(m.dbStorage.Start(m.ctx))
	goutil.PanicUnhandledError(m.peerConnector.Start(m.ctx))
	goutil.PanicUnhandledError(m.dbServer.Start())
	return nil
}
