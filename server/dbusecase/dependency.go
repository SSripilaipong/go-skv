package dbusecase

type Dependency struct {
	storageChan chan<- any
}

func NewDependency(storageChan chan<- any) *Dependency {
	return &Dependency{
		storageChan: storageChan,
	}
}
