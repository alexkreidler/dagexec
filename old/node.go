package old

import (
	"github.com/goombaio/dag"
	"github.com/satori/go.uuid"
)

// Node represents a computational task to be run
type Node struct {
	Name string
	// the Node id matches the DAG id
	id    string
	v     *dag.Vertex
	graph *Graph
	Executor
	// done represents if this node has been marked as completed
	done bool
	// err represents an error that occurred in the parent graph of the node
	errs []error
}

// EasyGraphNode is an interface for easily manipulating GoExecutor type graphs
type EasyGraphNode interface {
	// A Child represents an immediate child of the node
	// The returned node is exactly one step further down the graph
	AddChild(childNode *Node) *Node
	// Children represents a step of nested/sequential Nodes. Returns the furthest child node
	Children(listOfChildren ...*Node) *Node

	// Adds another parallel node to a parent node. Returns the parent node
	AddParallel(task *Node) *Node
	// Parallel adds a set of  parallel nodes to a parent node. Returns the parent node
	Parallel(tasks ...*Node) *Node

	// Done completes a section of a graph. It will return the errors accumulated from any previous graph operations
	Done() []error

	// End marks the node as the end node, which makes the graph run it.
	// Only one node can be the end node, and the call will fail if another end node is already set
	End() error

	//TODO: implement merge
	//Merge takes a bunch of split up nodes (e.g. the result from parallel), and then combines their data somehow and calls the specified function
	//This is needed to get a final singular result from the graph
	//Merge(...Node)
}

func newNode(name string, executor Executor) *Node {
	id := uuid.NewV4()
	v := dag.NewVertex(id.String(), nil)
	n := Node{
		Name:     name,
		id:       id.String(),
		v:        v,
		Executor: executor,
	}
	return &n
}
