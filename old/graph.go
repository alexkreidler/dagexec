package old

import (
	"fmt"
	"github.com/goombaio/dag"
)

// A graph represents a (limited) computational graph
// In our case, it must start with one singular node and end with another singular node
type Graph struct {
	Name     string
	RootNode *Node
	EndNode  *Node
	dag      *dag.DAG
}

func (g *Graph) NewNode(name string, f GenericFunction) (*Node, error) {
	n := newNode(name, GoExecutor{Function: f})
	err := g.dag.AddVertex(n.v)
	if err != nil {
		return nil, err
	}
	n.graph = g
	return n, nil
}

func (g Graph) Execute() (interface{}, error) {
	if g.EndNode == nil {
		return nil, fmt.Errorf("Failed to execute: no EndNode set. Have you run node.End()?")
	}
	//TODO: traverse up from the endnode and only run the relevant nodes
	//panic("implement me")
	return nil, nil
}

type EasyGraphGraph interface {
	// Synchronous execution of the graph all the way to the end node, returns the final result
	Execute() (interface{}, error)
	NewNode(string, GenericFunction) *Node
}

func NewGraph(name string) Graph {
	g := Graph{
		Name:     name,
		RootNode: newNode(name+" Root Node", nil),
		EndNode:  nil,
		dag:      dag.NewDAG(),
	}
	err := g.dag.AddVertex(g.RootNode.v)
	if err != nil {
		return Graph{}
	}
	g.RootNode.graph = &g
	return g
}
