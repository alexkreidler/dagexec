package old

import "fmt"

func (n *Node) AddChild(childNode *Node) *Node {
	err := ensure(childNode)
	if err != nil {
		n.errs = append(n.errs, fmt.Errorf("invalid child node %s: %s", childNode.id, err.Error()))
	}

	err = n.graph.dag.AddEdge(n.v, childNode.v)
	if err != nil {
		n.errs = append(n.errs, err)
	}
	return childNode
}

func (n *Node) Children(listOfChildren ...*Node) *Node {
	prevNode := n
	for _, child := range listOfChildren {
		prevNode = prevNode.AddChild(child)
	}
	return prevNode
}

func (n *Node) AddParallel(task *Node) *Node {
	err := ensure(task)
	if err != nil {
		n.errs = append(n.errs, fmt.Errorf("invalid parallel child/task node %s: %s", task.id, err.Error()))
	}

	err = n.graph.dag.AddEdge(n.v, task.v)
	if err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

func (n *Node) Parallel(tasks ...*Node) *Node {
	newNode := n
	for _, task := range tasks {
		newNode = n.AddParallel(task)
	}
	return newNode
}

func (n *Node) Done() []error {
	if n.done {
		return []error{fmt.Errorf("node %s already marked as done", n.Name)}
	}
	n.done = true
	if len(n.errs) != 0 {
		return n.errs
	} else {
		return nil
	}
}

func (n *Node) End() error {
	if !n.done {
		return fmt.Errorf("node %s not marked as done", n.Name)
	}
	if n.graph.EndNode != nil {
		return fmt.Errorf("failed setting %s as end node, graph %s already has end node %s", n.Name, n.graph.Name, n.graph.EndNode.Name)
	} else {
		n.graph.EndNode = n
		return nil
	}
}

func ensure(n *Node) error {
	if n.id == "" || n.Name == "" {
		return fmt.Errorf("ID and Name must be set")
	}
	if n.v == nil {
		return fmt.Errorf("no vertex set for node")
	}

	if n.id != n.v.ID {
		return fmt.Errorf("node ID is not the same as vertex ID")
	}

	if n.graph == nil {
		return fmt.Errorf("no graph set for node")
	}

	_, err := n.graph.dag.GetVertex(n.id)
	if err != nil {
		return fmt.Errorf("DAG does not contain the node vertex")
	}

	return nil
}
