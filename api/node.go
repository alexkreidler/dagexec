package api

import (
	"github.com/alexkreidler/dagexec/lib/observer"
)

// Input represents the input to a function. It can be any generic data
//Todo: hasing/determinism to skip/cache functions based on input
type Input interface{}

type Output interface{}

// A function represents all the data required to execute a computational function
// The executor ID is matched with the executor of the task to make sure the function can be executed properly in the given environment
type Function interface {
	ExecutorID() ExecutorID
	Data() interface{}
}

type ExecutorID string

type Executor interface {
	// Must return a unique name
	ID() ExecutorID

	Run(task Task)

	// Has to implement observer notifier to notify the T
	observer.Notifier
}

type Task struct {
	Function
	State

	trials       []*Trial
	currentTrial *Trial
}

type Trial struct {
	t *Task
	Input
	state State
	exec  *Executor
}
