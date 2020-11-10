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
	//				     0
	//				  /     \
	//				1	 2
	// 			       / \      / \
	// 			      3   4    5   6

	nodes := []*Node{root, node1, node2, node3, node4, node5, node6}
	var DAG = createDAG(nodes)
	var LCA, _ = findLCA(DAG, node5, node6)
	if LCA.key != 2 {
		t.Errorf("The LCA of 5 and 6 should be 2")
	}
	LCA, _ = findLCA(DAG, node3, node4)
	if LCA.key != 1 {
		t.Errorf("The LCA of 3 and 4 should be 1")
	}
	LCA, _ = findLCA(DAG, node3, node6)
	if LCA.key != 0 {
		t.Errorf("The LCA of 3 and 6 should be 0")
	}
	LCA, _ = findLCA(DAG, node4, node5)
	if LCA.key != 0 {
		t.Errorf("The LCA of 4 and 5 should be 0")
	}
	LCA, _ = findLCA(DAG, node1, node6)
	if LCA.key != 0 {
		t.Errorf("The LCA of 1 and 6 should be 0")
	}
}

func TestFindLCA_nil (t *testing.T) {
	var root = createNode(0, nil)
	nodes:= []*Node{root}
	var DAG = createDAG(nodes)
	//       0
	//    /     \
	//  nil     nil
	var LCA, _ = findLCA(DAG, nil, nil)
	if LCA != nil {
		t.Errorf("LCA with Nil nodes should return Nil")
	}
	DAG.nodes[0] = nil
	//      nil
	//    /     \
	//  nil     nil
	LCA, _ = findLCA(DAG, nil, nil)
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
	//				     0
	//				  /     \
	//				1	 2
	// 			       / \      / \
	// 			      3   4    5   6
	// Node 7 is not in the tree

	nodes := []*Node{root, node1, node2, node3, node4, node5, node6}
	var DAG = createDAG(nodes)
	var LCA, _ = findLCA(DAG, node5, node7)
	if LCA != nil {
		t.Errorf("The LCA should ne Nil as 7 is not in the tree")
	}

}

func TestDag (t *testing.T) {
	var node8  = createNode (8, nil)
	var node7  = createNode (7, []*Node{node8})
	var node6  = createNode (6, []*Node{node8})
	var node5  = createNode (5, nil)
	var node4  = createNode (4, nil)
	var node3  = createNode (3, []*Node{node5, node6})
	var node2  = createNode (2, nil)
	var node1  = createNode (1, []*Node{node3, node4})
	var node0  = createNode (0, []*Node{node2, node3})

	nodes := []*Node{node0, node1, node2, node3, node4, node5, node6, node7, node8}
	var DAG = createDAG(nodes)
	//					  0      1
	//				       /     \/    \
	//				      2	      3     4
	// 			       		    /   \
	// 			  	           5     6   7
	//						  \ /
	//					           8
	var LCA, _ = findLCA(DAG, node5, node6)
	if LCA.key != 3 {
		t.Errorf("The LCA of 5 and 6 should be 3")
	}
	LCA, _ = findLCA(DAG, node3, node4)
	if LCA.key != 1 {
		t.Errorf("The LCA of 3 and 4 should be 1")
	}
	LCA, _ = findLCA(DAG, node7, node8)
	if LCA != nil {
		t.Errorf("The LCA of 7 and 8 should be nil")
	}
	LCA, _ = findLCA(DAG, node2, node4)
	if LCA != nil {
		t.Errorf("The LCA of 2 and 4 should be nil")
	}

}
func TestNoConnection(t *testing.T){

	var node4  = createNode (4, nil)
	var node3  = createNode (3, nil)
	var node2  = createNode (2, []*Node{node3})
	var node1  = createNode (1, []*Node{node3})
	var node0  = createNode (0, []*Node{node1, node2})
	nodes := []*Node{node0, node1, node2, node3, node4}
	var DAG = createDAG(nodes)
	//       node3        node2
	//      /     \
	//   node4   node5
	//      \     /
	//       node1
	var LCA, _ = findLCA(DAG, node3, node4)

	if LCA != nil {
		t.Errorf("The LCA of 3 and 4 should be nil")
	}
}
