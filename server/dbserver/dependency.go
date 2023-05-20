package dbserver

import "go-skv/server/dbusecase"

type Dependency struct {
	GetValueUsecase dbusecase.GetValueFunc
}
