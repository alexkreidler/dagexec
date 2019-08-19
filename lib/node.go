package lib

import (
	"fmt"
	"github.com/goombaio/dag"
	"github.com/satori/go.uuid"
)


// Node represents a computational task to be run
type Node struct {
	Name string
	// the Node id matches the DAG id
	id string
	v *dag.Vertex
	graph *Graph
	Executor
}

func ensure(n *Node) error {
	if (n.v == nil) {
		return fmt.Errorf("No vertex set for node")
	}

	if n.id != n.v.ID {
		return fmt.Errorf("Node ID is not the same as vertex ID")
	}

	if (n.graph == nil) {
		return fmt.Errorf("No graph set for node")
	}

	_, err := n.graph.dag.GetVertex(n.id)
	if err != nil {
		return fmt.Errorf("DAG does not contain the node vertex")
	}

	return nil
}

func (n Node) Child(cn Node) error {
	e := ensure(&cn)
	if e != nil {
		return fmt.Errorf("Invalid child node %s", cn.id)
	}
	//n.graph.dag.Add
}

func (n Node) Children(...Node) {
	panic("implement me")
}

func (n Node) Parallel(...Node) {
	panic("implement me")
}

type EasyGraphNode interface {
	// A Child represents an immediate child of the node
	// The returned node is exactly one step further down the graph
	Child(Node)
	// Children represents a step of nested/sequential Nodes
	Children(...Node)
	// Parallel represents a set of nodes which can be executed in parallel
	Parallel(...Node)

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

//func NewNode(name string, f GenericFunction) *Node {
//	return newNode(name,  GoExecutor{Function:f})
//}
