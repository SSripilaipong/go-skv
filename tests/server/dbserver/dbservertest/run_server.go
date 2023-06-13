package dbservertest

import (
	"fmt"
	"go-skv/server/dbserver"
	"go-skv/server/dbusecase"
)

func RunWithGetValueUsecase(usecase dbusecase.Interface, execute func(server dbserver.Interface) error) error {
	return runWithUsecases(usecase, execute)
}

func RunWithSetValueUsecase(usecase dbusecase.Interface, execute func(server dbserver.Interface) error) error {
	return runWithUsecases(usecase, execute)
}

func runWithUsecases(usecase dbusecase.Interface, execute func(server dbserver.Interface) error) error {
	server := dbserver.New(0, usecase)
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
