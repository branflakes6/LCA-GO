package main

import "errors"

type Node struct  {
   key int
   children []*Node
}

func createNode(key int, children []*Node) *Node {
	var node Node
	node.key = key
	node.children = children
	return &node
}

func findLCA (root, a, b *Node) (*Node, error) {
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
		var temp, _ = findLCA( root.children[i], a, b)
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