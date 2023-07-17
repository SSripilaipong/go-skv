package replicaupdatertest

import (
	"context"
	"go-skv/common/actormodel"
)

type RecordUpdaterFactoryMock struct {
	New_Return actormodel.ActorRef
}

func (r *RecordUpdaterFactoryMock) New(ctx context.Context, key string, value string) actormodel.ActorRef {
	return r.New_Return
}
