package dbmanager

import (
	"context"
	"go-skv/util/goutil"
)

func (m *manager) Start() error {
	goutil.PanicUnhandledError(m.dbStorage.Start())
	goutil.PanicUnhandledError(m.peerConnector.Start(context.Background()))
	goutil.PanicUnhandledError(m.dbServer.Start())
	return nil
}
