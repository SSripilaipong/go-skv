package dbmanager

import (
	"go-skv/util/goutil"
)

func (m *manager) Start() error {
	goutil.PanicUnhandledError(m.dbStorage.Start())
	goutil.PanicUnhandledError(m.peerConnector.Start())
	goutil.PanicUnhandledError(m.dbServer.Start())
	return nil
}
