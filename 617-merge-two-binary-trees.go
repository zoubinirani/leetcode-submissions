// 617. Merge Two Binary Trees
// Leetcode Easy done in Golang
// Runtime: 28 ms (faster than 90%)
// Memory Usage: 8.8 MB (less than 100%)

func traverse(t1, t2 *TreeNode) {
	t1.Val += t2.Val
	
	if (t1.Left == nil) && (t2.Left != nil) {
			var tree TreeNode = TreeNode{0, nil, nil}
			t1.Left = &tree
			traverse(t1.Left, t2.Left)
	} else if (t2.Left != nil) {
			traverse(t1.Left, t2.Left)
	}
	
	if (t1.Right == nil) && (t2.Right != nil) {
			var tree TreeNode = TreeNode{0, nil, nil}
			t1.Right = &tree
			traverse(t1.Right, t2.Right)
	} else if (t2.Right != nil) {
			traverse(t1.Right, t2.Right)
	}    
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	
	if (t1 == nil) {
			return t2
	}
	
	if (t2 == nil) {
			return t1
	}
	
	traverse(t1, t2)
	return t1
}