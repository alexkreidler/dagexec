package old

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNodes(t *testing.T) {

	Convey("with a graph", t, func() {

		g := NewGraph("test1")
		Convey("create node", func() {

			f := func(interface{}) interface{} {
				return "hello world"
			}

			n, err := g.NewNode("hello world", f)

			So(err, ShouldBeNil)

			Convey("get node", func() {
				v, err := g.dag.GetVertex(n.id)
				So(err, ShouldBeNil)
				So(v, ShouldNotBeZeroValue)
			})

			Convey("remove node", nil)

			Convey("change node", nil)

			Convey("run node", nil)

		})

	})

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
	for i, f := range fs {
		name := o[i]
		s := f(nil)
		t.Logf("Got %s for item %s", s, name)
	}

	graph := NewGraph("testGraph")
	g := &graph

	t.Log("creating computational graph")
	prevNode := g.RootNode
	if prevNode == nil {
		t.Error("failed to create root node")
	}
	t.Log(prevNode)
	for i, f := range fs {
		name := o[i]
		n, err := g.NewNode(name, f)
		if err != nil {
			t.Error("failed creating node:", err)
		}
		//t.Logf("%#v", graph.dag)
		if prevNode == nil {
			t.Error("previous Node is nil")
		} else {
			prevNode = prevNode.AddChild(n)
		}
	}
	if prevNode == nil {
		t.Error("previous Node is nil")
	} else {
		errs := prevNode.Done()
		if len(errs) != 0 {
			t.Error(errs)
		}
		err := prevNode.End()
		if err != nil {
			t.Error(err)
		}
	}
	t.Logf("%#v", prevNode)
	t.Logf("%#v", prevNode.graph)
	t.Logf("%#v", g)

	//t.Ass
	//Now prevNode is the endNode of the graph

	data, err := prevNode.graph.Execute()
	if err != nil {
		t.Error(err)
	}
	t.Log(data)

	//root :=
	//t.Log("Got: ", s)
	//
	//n := NewNode("hello world", f)
	//t.Logf("%#v", n)
}

// Example routine to create: template all files, then execute files in specific order

func templateService(svc string) GenericFunction {
	return func(interface{}) interface{} {
		// TODO: create the actual service
		return svc
	}
}

func createService(svc string) GenericFunction {
	return func(interface{}) interface{} {
		// TODO: create the actual service
		return svc
	}
}

func TestEasyGraph(t *testing.T) {

	graph := NewGraph("testGraph")

	services := []string{"catalog", "edge", "control"}

	for _, service := range services {
		f := templateService(service)
		n, _ := graph.NewNode("template service: "+service, f)
		c := graph.RootNode.AddParallel(n)
		f = createService(service)
		b, _ := graph.NewNode("create service: "+service, f)
		c.AddChild(b)
	}

}

//func TestSuperEasyGraph(t (testin))

func TestSuperEasyGraph(t *testing.T) {
	graph := NewGraph("graph")
	// Easily create child functions that run in sequential order
	graph.RootNode.Child(func() {
		fmt.Println("child 1")
	}).Child(func() {
		fmt.Print("child 2")
		//	And parallel child functions which run in parallel
	}).Pl(func() {
		fmt.Println("downloading files")
	}, func() {
		fmt.Println("downloading f2")
		//	Then waits before completing
	}).Child(func() {
		fmt.Println("done")
	})
}
