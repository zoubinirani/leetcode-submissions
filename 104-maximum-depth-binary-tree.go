// 104. Maximum Depth of Binary Tree
// Leetcode Easy done in Golang
// Runtime 4ms (faster than 92.6%)
// Memory usage: 4.9 MB (less than 83.3%)

type depthValues []int

func (d *depthValues) appToEnd(depth int) {
    *d = append(*d, depth)
}

func traverse(head *TreeNode, depth int, depthValPointer *depthValues) {
    if head == nil {
        depthValPointer.appToEnd(depth)
        return
    }
    traverse(head.Left, depth + 1, depthValPointer)
    traverse(head.Right, depth + 1, depthValPointer)
}

func maxDepth(root *TreeNode) int {
    depthVal:= depthValues{}
    var depthValPointer *depthValues = &depthVal
    
    traverse(root, 0, depthValPointer)
    
    var maxVal int 
    for _, val := range(depthVal) {
        if (val > maxVal) {
            maxVal = val
        }
    }
        
    return maxVal
}