package replicaupdater

import (
	"context"
	"go-skv/common/actormodel"
)

type RecordUpdaterFactory interface {
	New(ctx context.Context, key string, value string) actormodel.ActorRef
}
