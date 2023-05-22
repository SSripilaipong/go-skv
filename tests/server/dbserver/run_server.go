package dbserverTest

import (
	"fmt"
	"go-skv/server/dbmanager"
	"go-skv/server/dbserver"
	"go-skv/server/dbusecase"
)

func RunWithPortAndGetValueUsecase(port int, usecase dbusecase.GetValueFunc, execute func(server dbmanager.DbServer) error) error {
	return runWithPortAndDependency(port, dbserver.Dependency{
		GetValueUsecase: usecase,
	}, execute)
}

func RunWithPortAndSetValueUsecase(port int, usecase dbusecase.SetValueFunc, execute func(server dbmanager.DbServer) error) error {
	return runWithPortAndDependency(port, dbserver.Dependency{
		SetValueUsecase: usecase,
	}, execute)
}

func runWithPortAndDependency(port int, dependency dbserver.Dependency, execute func(server dbmanager.DbServer) error) error {
	server := dbserver.New(port, dependency)
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
