package dbusecase

import "go-skv/server/storage"

type Dependency struct {
	storageChan chan storage.Packet
}

func NewDependency(storageChan chan storage.Packet) *Dependency {
	return &Dependency{
		storageChan: storageChan,
	}
}
