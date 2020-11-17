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
func deleteNode(root *TreeNode , key int)*TreeNode{

	// 删除节点分为三种情况：
	// 1、只有左节点 替换为右
	// 2、只有右节点 替换为左
	// 3、有左右子节点 左子节点连接到右边最左节点即可

	if root == null {
		return root
	}

	if root.Val < key {
		//search
		root.Right = deleteNode(root.Right , key)
	}else if root.Val > key{
		// search
		root.Left = deleteNode(root.Left , key)
	}else if root.Val == key{
		if root.Left == nil {
			return root.Right
		}else if root.Right == nil {
			return root.Left
		}else {
			cur := root.Right
			for cur.Left != nil  {
				cur = cur.Left
			}

			cur.Left = root.Left

			return root.Right
		}
	}

	return root

}

