package main

import "testing"


func TestFindLCA(t *testing.T) {
	var node6 = createNode(6, nil)
	var node5 = createNode(5, nil)
	var node4 = createNode(4, nil)
	var node3 = createNode(3, nil)
	var node2 = createNode(2, []*Node{node5, node6})
	var node1 = createNode(1, []*Node{node3, node4})
	var root = createNode(0, []*Node{node1, node2})

	var LCA, _ = findLCA(root, node5, node6)
	if LCA.key != 2 {
		t.Errorf("The LCA of 5 and 6 should be 2")
	}
	LCA, _ = findLCA(root, node3, node4)
	if LCA.key != 1 {
		t.Errorf("The LCA of 3 and 4 should be 1")
	}
	LCA, _ = findLCA(root, node3, node6)
	if LCA.key != 0 {
		t.Errorf("The LCA of 3 and 6 should be 0")
	}
	LCA, _ = findLCA(root, node4, node5)
	if LCA.key != 0 {
		t.Errorf("The LCA of 4 and 5 should be 0")
	}
	LCA, _ = findLCA(root, node1, node6)
	if LCA.key != 0 {
		t.Errorf("The LCA of 1 and 6 should be 0")
	}
}

func TestFindLCA_nil (t *testing.T) {
	var root = createNode(0, nil)
	var LCA, _ = findLCA(root, nil, nil)
	if LCA != nil {
		t.Errorf("LCA with Nil nodes should return Nil")
	}
	root = nil
	LCA, _ = findLCA(root, nil, nil)
	if LCA != nil {
		t.Errorf("LCA with Nil nodes should return Nil")
	}
}
func TestNotInTree (t *testing.T) {
	var node7 = createNode(7, nil)
	var node6 = createNode(6, nil)
	var node5 = createNode(5, nil)
	var node4 = createNode(4, nil)
	var node3 = createNode(3, nil)
	var node2 = createNode(2, []*Node{node5, node6})
	var node1 = createNode(1, []*Node{node3, node4})
	var root = createNode(0, []*Node{node1, node2})

	var LCA, _ = findLCA(root, node5, node7)
	if LCA != nil {
		t.Errorf("The LCA should ne Nil as 7 is not in the tree")
	}

}