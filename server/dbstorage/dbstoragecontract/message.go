package dbstoragecontract

import "go-skv/common/actormodel"

type GetRecord struct {
	Key     string
	ReplyTo actormodel.ActorRef
}
