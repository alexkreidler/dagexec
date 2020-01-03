package node

import (
	"github.com/alexkreidler/dagexec/newlib/api"
	"github.com/alexkreidler/dagexec/newlib/observer"
)

// A Node represents a single computational task to be run. It has a 1-1 relationship with a Task
// For future, might have multiple tasks in a single node (setup/teardown)
type Node struct {
	id          int64
	Name        string
	DisplayName string
	Task
}

// If a task fails, it can be retried if need be
//
type Task struct {
	// The executor to run the task
	Runner *api.Executor

	// The subscriber for the state of the task
	StateObserver observer.Observer

	State api.State
}
