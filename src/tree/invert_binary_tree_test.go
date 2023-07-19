package tree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	// Test case 1: A simple binary tree to invert.
	//     1           1
	//    / \   =>    / \
	//   2   3       3   2
	//  / \             / \
	// 4   5           5   4
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	expected1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 4},
		},
	}

	result1 := invertTree(root1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected: %v, got: %v", expected1, result1)
	}

	// Test case 2: An empty tree (nil root).
	root2 := (*TreeNode)(nil)
	expected2 := (*TreeNode)(nil)

	result2 := invertTree(root2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected: %v, got: %v", expected2, result2)
	}

	// Add more test cases as needed.
}
