package tree

import (
    "fmt"
    "os"
)

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func MaxDepth(root *TreeNode) int {

    node := []interface{}{3, 9, 20, 15, 7}

    data := GetTree(root, node...)
    showTree(data)
    return 1
}

//生成二叉树
func GetTree(treeNode *TreeNode, data ...interface{}) *TreeNode {
    //生成二叉树
    for _, v := range data {
        fmt.Println(v)
        //treeNode = SaveSortTree(treeNode, v)
        treeNode = saveTree(treeNode, v)
    }

    return treeNode
}

//保存树结构 有序二叉树
func SaveSortTree(root *TreeNode, data interface{}) *TreeNode {
    //
    if root == nil {
        root = &TreeNode{}
    }
    if root.Val == 0 {
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

//保存二叉树 无序
func saveTree(root *TreeNode, data ...interface{}) *TreeNode {
    num := 1
    fmt.Println(data[0:num])
    os.Exit(1)
    if data == nil {
        return nil
    }

    if root == nil {
        root = &TreeNode{}
    }
    if root.Val == 0 {
        root.Val = data[0].(int)
        return root
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
    fmt.Println(root.Val)
    showTree(root.Left)
    showTree(root.Right)
    return
}
