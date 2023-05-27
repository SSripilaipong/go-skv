package stop

import (
	"github.com/stretchr/testify/assert"
	dbmanagerTest "go-skv/tests/server/dbmanager"
	"testing"
)

func Test_should_stop_db_server(t *testing.T) {
	dbServer := &dbmanagerTest.DbServerMock{}
	mgr := dbmanagerTest.NewWithDbServer(dbServer)

	_ = doStop(mgr)

	assert.True(t, dbServer.Stop_IsCalled)
}

func Test_should_stop_db_storage(t *testing.T) {
	dbStorage := &dbmanagerTest.DbStorageMock{}
	mgr := dbmanagerTest.NewWithDbStorage(dbStorage)

	_ = doStop(mgr)

	assert.True(t, dbStorage.Stop_IsCalled)
}
