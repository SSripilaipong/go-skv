package dbserver

import (
	"go-skv/server/dbserver/dbservercontroller"
	"go-skv/server/dbserver/dbusecase"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func New(port int, repo dbstoragecontract.Storage) Interface {
	return dbservercontroller.New(port, dbusecase.New(repo))
}
