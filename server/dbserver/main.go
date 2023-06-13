package dbserver

import (
	"go-skv/server/dbserver/dbservercontroller"
	"go-skv/server/dbserver/dbusecase"
	"go-skv/server/dbstorage"
)

func New(port int, repo dbstorage.RepositoryInteractor) Interface {
	return dbservercontroller.New(port, dbusecase.New(repo))
}
