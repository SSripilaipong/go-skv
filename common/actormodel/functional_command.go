package actormodel

func Do[S any](execute func(state *S)) Command[S] {
	return functionalCommand[S]{execute}
}

type functionalCommand[S any] struct {
	execute func(state *S)
}

func (c functionalCommand[S]) Execute(state *S) {
	c.execute(state)
}
