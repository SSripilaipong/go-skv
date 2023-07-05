package dbmanager

import (
	"context"
	"go-skv/common/util/goutil"
)

func (m manager) Start() error {
	goutil.PanicUnhandledError(m.dbStorage.Start(context.Background()))
	goutil.PanicUnhandledError(m.peerConnector.Start(m.ctx))
	goutil.PanicUnhandledError(m.dbServer.Start())
	return nil
}
