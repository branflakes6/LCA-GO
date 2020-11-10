package main
import "errors"

type Node struct  {
   key int
   children []*Node
   visited bool
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
	node.visited = false
	return &node
}

func findLCA (DAG *DAG, a, b *Node) (*Node, error) {

	for i := 0; i < len(DAG.nodes); i++{
		var LCA *Node = nil
		LCA, _ = searchGraph(DAG.nodes[i], a, b)
		if LCA != nil {
			if a.visited == true && b.visited == true {
				resetVisited(DAG)
				return LCA, nil
			}
		}
		resetVisited(DAG)
	}
	resetVisited(DAG)
	return nil, errors.New("one or more nodes is not in the graph")
}

func resetVisited (DAG *DAG) () {
    for i := 0; i < len(DAG.nodes); i++{
    	if DAG.nodes[i] != nil {
			DAG.nodes[i].visited = false
		}
    }
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
			var temp *Node
			temp, _ = searchGraph(root.children[i], a, b)
			if temp != nil {
				count++
				child = temp
				child.visited = true
			}
			if count >= 2 {
				return root, nil
			}
		}
	}
		return child, nil
}