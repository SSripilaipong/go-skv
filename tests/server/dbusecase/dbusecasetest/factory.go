package dbusecasetest

import (
	"go-skv/server/dbstorage"
	"go-skv/server/dbusecase"
)

func NewUsecaseWithRepo(repo dbstorage.RepositoryInteractor) dbusecase.Interface {
	return dbusecase.New(repo)
}
