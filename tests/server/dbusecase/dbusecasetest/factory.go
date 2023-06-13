package dbusecasetest

import (
	"go-skv/server/dbserver/dbusecase"
	"go-skv/server/dbstorage"
)

func NewUsecaseWithRepo(repo dbstorage.RepositoryInteractor) dbusecase.Interface {
	return dbusecase.New(repo)
}
