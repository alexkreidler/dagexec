package lib

import (
	"fmt"
	"github.com/goombaio/dag"
)

// A graph represents a (limited) computational graph
// In our case, it must start with one singular node and end with another singular node
type Graph struct {
	Name string
	RootNode *Node
	EndNode *Node
	dag *dag.DAG
}

func (g *Graph) NewNode(name string, f GenericFunction) *Node {
	n := newNode(name,  GoExecutor{Function:f})
	g.dag.AddVertex(n.v)
	n.graph = g
	return n
}

func (g Graph) Execute() error {
	if g.EndNode == nil {
		return fmt.Errorf("Failed to execute: no EndNode set. Have you run node.End()?")
	}
	//TODO: traverse up from the endnode and only run the relevant nodes
	//panic("implement me")
}

type EasyGraphGraph interface {
	// Synchronous execution of the graph all the way to the end node
	Execute() error
	NewNode(string, GenericFunction) *Node
}

func NewGraph(name string) Graph {
	return Graph{
		Name:     name,
		RootNode: newNode(name + " Root Node", nil),
		EndNode:  nil,
		g: dag.NewDAG(),
	}
}