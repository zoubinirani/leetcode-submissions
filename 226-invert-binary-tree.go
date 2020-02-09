// 226. Invert Binary Tree
// Leetcode Easy done in Golang
// Runtime <1ms (faster than 100%)
// Memory usage: 2.2 MB (less than 100%)

func inverse(treePointer *TreeNode, root *TreeNode) {
     if root.Left != nil {
        var leftTree = TreeNode{root.Left.Val, nil, nil}
        treePointer.Right = &leftTree
        inverse(treePointer.Right, root.Left)        
    } 
    
    if root.Right != nil {
        var rightTree = TreeNode{root.Right.Val, nil, nil}
        treePointer.Left = &rightTree
        inverse(treePointer.Left, root.Right)
    }
}

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var newTree = TreeNode{root.Val, nil, nil}
    var newTreePointer *TreeNode = &newTree
    inverse(newTreePointer, root)
    return newTreePointer
}
