package dbserverTest

import (
	"fmt"
	"go-skv/server/dbmanager"
	"go-skv/server/dbserver"
	"go-skv/server/dbusecase"
)

func RunWithPortAndGetValueUsecase(port int, usecase dbusecase.GetValueFunc, execute func(server dbmanager.DbServer) error) error {
	server := dbserver.New(port, dbserver.Dependency{
		GetValueUsecase: usecase,
	})
	if err := server.Start(); err != nil {
		panic(fmt.Errorf("unexpected error"))
	}
	defer func() {
		if err := server.Stop(); err != nil {
			panic(fmt.Errorf("unexpected error"))
		}
	}()
	return execute(server)
}
