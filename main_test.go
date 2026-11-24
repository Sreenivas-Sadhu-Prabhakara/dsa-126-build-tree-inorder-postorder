package main

import (
	"reflect"
	"testing"
)

func inorderOf(n *TreeNode) []int {
	if n == nil {
		return nil
	}
	return append(append(inorderOf(n.Left), n.Val), inorderOf(n.Right)...)
}

func postorderOf(n *TreeNode) []int {
	if n == nil {
		return nil
	}
	out := postorderOf(n.Left)
	out = append(out, postorderOf(n.Right)...)
	return append(out, n.Val)
}

func TestBuild(t *testing.T) {
	in := []int{9, 3, 15, 20, 7}
	post := []int{9, 15, 7, 20, 3}
	root := Build(in, post)
	if got := inorderOf(root); !reflect.DeepEqual(got, in) {
		t.Errorf("inorder = %v, want %v", got, in)
	}
	if got := postorderOf(root); !reflect.DeepEqual(got, post) {
		t.Errorf("postorder = %v, want %v", got, post)
	}

	if Build(nil, nil) != nil {
		t.Error("empty build should be nil")
	}
	// Right-skewed: inorder 1 2 3, postorder 3 2 1.
	skew := Build([]int{1, 2, 3}, []int{3, 2, 1})
	if skew.Val != 1 || skew.Right.Val != 2 || skew.Right.Right.Val != 3 || skew.Left != nil {
		t.Error("right-skewed build wrong")
	}
}
