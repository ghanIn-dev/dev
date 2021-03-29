package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	//s의 모든 노드를 방문하면서 compareTree호출해서
	//결과가 true면 true를 반환한다
	return DFSFunc(s, t, compareTree)
}

func DFSFunc(s, t *TreeNode, f func(*TreeNode, *TreeNode) bool) bool {
	if s == nil {
		if t == nil {
			return true
		}
		return false
	}

	if f(s, t) == true {
		return true
	}

	if DFSFunc(s.Left, t, f) == true {
		return true
	}

	return DFSFunc(s.Right, t, f)
}

func compareTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		if t2 == nil {
			return true
		}
		return false
	} else if t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if !compareTree(t1.Left, t2.Left) {
		return false
	}
	return compareTree(t1.Right, t2.Right)
}
