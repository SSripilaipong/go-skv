package dbserverTest

import (
	"fmt"
	"go-skv/server/dbserver"
	"go-skv/server/dbusecase"
)

func RunWithGetValueUsecase(usecase dbusecase.GetValueFunc, execute func(server dbserver.Interface) error) error {
	return runWithDependency(dbserver.Dependency{
		GetValueUsecase: usecase,
	}, execute)
}

func RunWithSetValueUsecase(usecase dbusecase.SetValueFunc, execute func(server dbserver.Interface) error) error {
	return runWithDependency(dbserver.Dependency{
		SetValueUsecase: usecase,
	}, execute)
}

func runWithDependency(dependency dbserver.Dependency, execute func(server dbserver.Interface) error) error {
	server := dbserver.New(0, dependency)
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
