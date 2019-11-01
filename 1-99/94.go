/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    
    if root == nil {
        return res
    }
    
    arr := []*TreeNode{}
    
    JLoop:
    for {
        JLoop1:
        for {
            if root.Left != nil {
                arr = append(arr, root)
                root = root.Left
            } else {
                break JLoop1
            }
        }
    
        res = append(res, root.Val)
        
        if root.Right != nil {
            root = root.Right            
        }else{
            JLoop2:
            for {
                if len(arr) > 0 {
                    root = arr[len(arr)-1]
                    arr = arr[:len(arr)-1]
                    res = append(res, root.Val)
                    if root.Right != nil {
                        root = root.Right
                        break JLoop2
                    }   
                } else {
                    break JLoop
                }
            }
        }
    }
    
    return res
}

func inorderTraversal2(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    
    if root.Left != nil {
        res = inorderTraversal(root.Left)    
    }
    
    res = append(res, root.Val)
    
    if root.Right != nil {
        t := inorderTraversal(root.Right)    
        res = append(res, t...)
    }
    
    return res
}