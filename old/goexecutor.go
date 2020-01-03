package old

type GenericFunction func(interface{}) interface{}

// A GoExecutor represents an Executor backed by a goroutine/go function
type GoExecutor struct {
	// we accept functions that can input and return generic values
	// use a stop channel to stop long running work?
	Function GenericFunction
	State
}

func (g GoExecutor) Configure() {
	panic("implement me")
}

func (g GoExecutor) Start() error {
	panic("implement me")
}

func (g GoExecutor) Stop() error {
	panic("implement me")
}
