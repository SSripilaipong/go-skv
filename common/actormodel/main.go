package actormodel

type Command[S any] interface {
	Execute(*S)
}
