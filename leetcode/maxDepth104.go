package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//topvoted
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

var md int

//tucker
func maxDepthT(root *TreeNode) int {
	if root == nil {
		return 0
	}

	md = 1

	if root.Left != nil {
		dfs(root.Left, 2)
	}

	if root.Right != nil {
		dfs(root.Right, 2)
	}

	return md
}

func dfs(node *TreeNode, depth int) {
	if depth > md {
		md = depth
	}

	if node.Left != nil {
		dfs(node.Left, depth+1)
	}

	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}
