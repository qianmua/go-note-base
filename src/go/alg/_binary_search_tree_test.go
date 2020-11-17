package alg

/**
 * @author
 * @date 2020/11/17 23:27
 * create by Goland
 * @version 1.0
 */ 
 
// 二叉树

// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。
//返回插入后二叉搜索树的根节点。 保证原始二叉搜索树中不存在新值。
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode , val int)*TreeNode{
	if root == nil {
		return &TreeNode{Val:val}
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right , val)
	}else {
		root.Left = insertIntoBST(root.Left , val)
	}

	return root

}

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的  key  对应的节点，并保证二叉搜索树的性质不变。
//返回二叉搜索树（有可能被更新）的根节点的引用。


