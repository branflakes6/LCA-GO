package main
import "errors"

type Node struct  {
   key int
   children []*Node
}
type DAG struct {
	nodes []*Node
}
func createDAG (nodes []*Node) *DAG {
	var dag DAG
	dag.nodes = nodes
	return &dag
}
func createNode(key int, children []*Node) *Node {
	var node Node
	node.key = key
	node.children = children
	return &node
}

func findLCA (root, a, b *Node) (*Node, error) {

	var LCA *Node = nil
	LCA, _ = searchGraph(root, a, b)
	if LCA != a && LCA != b {
		return LCA, nil

	}
	return nil, errors.New("one or more nodes is not in the graph")
}

func findLCADAG (DAG *DAG, a, b *Node) (*Node, error) {

	for i := 0; i < len(DAG.nodes); i++{
		var LCA *Node = nil
		LCA, _ = searchGraph(DAG.nodes[i], a, b)
		if LCA != a && LCA != b {
			return LCA, nil

		}
	}
	return nil, errors.New("one or more nodes is not in the graph")
}

func searchGraph (root, a, b *Node) (*Node, error) {

	if a == nil || b == nil {
		return nil, errors.New("cannot search for a Nil Node")
	}
	if root == nil || root.key == a.key || root.key == b.key {
		return root, nil
	}
	var child *Node = nil
	count := 0
	if root.children != nil {
		for i := 0; i < len(root.children); i++ {
		var temp, _ = searchGraph( root.children[i], a, b)
		if temp!= nil {
			count++
			child = temp
		}
		if count >= 2{
			return root, nil
		}
		}
	}
	return child, nil
}