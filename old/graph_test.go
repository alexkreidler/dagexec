package old

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGraph(t *testing.T) {

	Convey("create some graphs", t, func() {
		names := []string{"test1", "johnny", "blablabla"}
		for _, name := range names {
			graph := NewGraph(name)
			So(graph.Name, ShouldEqual, name)
			So(graph.RootNode.Name, ShouldEqual, name+" Root Node")
		}
	})

	Convey("create a graph", t, func() {
		graph := NewGraph("TEST1")
		Convey("using a basic function", func() {

			f := func(interface{}) interface{} {
				return "gotcha"
			}

			Convey("add a single node", func() {
				n, err := graph.NewNode("node1", f)

				So(err, ShouldBeNil)
				So(n, ShouldNotBeZeroValue)

				Convey("execute single node", func() {
					//res, err := graph.Execute()
					//So(err, ShouldBeNil)
					//So(res, ShouldNotBeZeroValue)
				})

			})

		})

	})

}
