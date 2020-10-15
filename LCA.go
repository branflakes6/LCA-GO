package main
import "fmt"

func main() {
	var node6 Node
	node6.key = 6
	node6.children = nil
	var node5 Node
	node5.key = 5
	node5.children = nil
	var node4 Node
	node4.key = 4
	node4.children = nil
	var node3 Node
	node3.key = 3
	node3.children = nil
	var node2 Node
	node2.key = 2
	node2.children = []*Node {&node5, &node6}
	var node1 Node
	node1.key = 1
	node1.children = []*Node {&node3, &node4}
	var root Node
	root.key = 0
	root.children = []*Node {&node1, &node2}
	var LCA = findLCA(&root, &node5, &node6)
	fmt.Println("The LCA of node 5 and 6 is ", LCA.key)
}

type Node struct  {
   key int
   children []*Node
}

func findLCA (root, a, b *Node) *Node {
	if root == nil || root.key == a.key || root.key == b.key {
		return root
	}
	var child *Node = nil
	count := 0
	if root.children != nil {
		for i := 0; i < len(root.children); i++ {
		var temp = findLCA( root.children[i], a, b)
		if temp!= nil {
			count++
			child = temp
		}
		if count >= 2{
			return root
		}
		}
	}
	return child
}