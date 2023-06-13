package dbservercontrollertest

import (
	"fmt"
	"go-skv/server/dbserver/dbservercontroller"
	"go-skv/server/dbserver/dbusecase"
)

func RunWithGetValueUsecase(usecase dbusecase.Interface, execute func(server dbservercontroller.Interface) error) error {
	return runWithUsecases(usecase, execute)
}

func RunWithSetValueUsecase(usecase dbusecase.Interface, execute func(server dbservercontroller.Interface) error) error {
	return runWithUsecases(usecase, execute)
}

func runWithUsecases(usecase dbusecase.Interface, execute func(server dbservercontroller.Interface) error) error {
	controller := dbservercontroller.New(0, usecase)
	if err := controller.Start(); err != nil {
		panic(fmt.Errorf("unexpected error"))
	}
	defer func() {
		if err := controller.Stop(); err != nil {
			panic(fmt.Errorf("unexpected error"))
		}
	}()
	return execute(controller)
}
