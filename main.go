package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Build reconstructs the unique binary tree from its inorder and postorder
// traversals (values distinct). Postorder read backwards hands out roots —
// so the right subtree must be built before the left.
func Build(inorder, postorder []int) *TreeNode {
	at := make(map[int]int, len(inorder))
	for i, v := range inorder {
		at[v] = i
	}
	post := len(postorder) - 1
	var build func(inLo, inHi int) *TreeNode
	build = func(inLo, inHi int) *TreeNode {
		if inLo >= inHi {
			return nil
		}
		v := postorder[post]
		post--
		n := &TreeNode{Val: v}
		mid := at[v]
		n.Right = build(mid+1, inHi) // right first: the cursor walks postorder backwards
		n.Left = build(inLo, mid)
		return n
	}
	return build(0, len(inorder))
}

func main() {
	root := Build([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}
