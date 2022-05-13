package tree

import (
    "fmt"
)

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func MaxDepth(root *TreeNode) int {

    node := []interface{}{3, 9, 20, nil, nil, 15, 7}

    data := GetTree(root, node...)
    showTree(data)

    //
    return 1
}

//生成二叉树
func GetTree(treeNode *TreeNode, data ...interface{}) *TreeNode {
    var mapTree =  map[int]*TreeNode{}
    //生成二叉树
    for k, v := range data {
        //顺序插入
        root := &TreeNode{}
        if v == nil {
            mapTree[k] = nil
        } else {
            root.Val = v.(int)
            mapTree[k] = root
        }
        //二叉搜索树
        //treeNode = SaveSortTree(treeNode, v)

    }
    nowTreeNode := &TreeNode{}
    nowTreeNode = mapTree[0]
    nowTreeNode.Left = mapTree[1]
    nowTreeNode.Right = mapTree[2]
    nowTreeNode.Left.Left = mapTree[3]
    nowTreeNode.Left.Right = mapTree[4]
    nowTreeNode.Right.Left = mapTree[5]
    nowTreeNode.Right.Right = mapTree[6]
    return nowTreeNode
}

//保存树结构 有序二叉树
func SaveSortTree(root *TreeNode, data interface{}) *TreeNode {
    //
    if root == nil {
        root = &TreeNode{}
    }
    if root.Val == 0 {
        if data == nil {
            return nil
        }
        root.Val = data.(int)
        return root
    }

    if data.(int) < root.Val {
        root.Left = SaveSortTree(root.Left, data)
    }

    if data.(int) > root.Val {
        root.Right = SaveSortTree(root.Right, data)
    }
    if data == 0 && root.Left == nil {
        root.Left = nil
    }
    if data == 0 && root.Right == nil {
        root.Right = nil
    }

    return root
}

/**
    前序遍历
 */
func showTree(root *TreeNode) {
    if root == nil {
        fmt.Println(nil)
        return
    }
    if root.Val == 0 {
        return
    }
    showTree(root.Left)
    showTree(root.Right)
    fmt.Println(root.Val)
    return
}

//计算深度

func light(root *TreeNode)  {
    
}
