package dbusecasetest

import (
	"go-skv/server/dbserver/dbusecase"
	"go-skv/server/dbstorage/dbstoragecontract"
)

func NewUsecaseWithRepo(repo dbstoragecontract.Storage) dbusecase.Interface {
	return dbusecase.New(repo)
}
