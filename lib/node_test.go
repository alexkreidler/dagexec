package lib

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	t.Log("Running testNode")
	f := func(interface{}) interface{} {
		return "hello world"
	}
	t.Log("testing function")
	s := f(nil)
	t.Log("Got: ", s)

	n := NewNode("hello world", f)
	t.Logf("%#v", n)
}


func TestGraphChildren(t *testing.T) {
	o := []string{"world", "bob", "johnny"}

	fs := make([]GenericFunction, 0)

	for _, name := range o {
		str := fmt.Sprintf("hello %s!", name)
		f := func(interface{}) interface{} {
			return str
		}
		fs = append(fs, f)
	}
	t.Log("testing functions")
	for _, f := range fs {
		s := f(nil)
		t.Log("Got: ", s)
	}

	graph := NewGraph("testGraph")

	t.Log("Creating computational graph")
	prevNode := graph.RootNode
	for i, f := range fs {
		name := o[i]
		prevNode.Child(NewNode(name, f))
	}
	prevNode.End()
	//Now prevNode is the endNode of the graph

	graph.Execute()

	//root :=
	//t.Log("Got: ", s)
	//
	//n := NewNode("hello world", f)
	//t.Logf("%#v", n)
}