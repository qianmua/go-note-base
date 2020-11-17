package alg



/**
 * @author
 * @date 2020/11/17 23:42
 * create by Goland
 * @version 1.0
 */

// 递归
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组  `char[]`  的形式给出。
func reverseString(s []byte){
	res := make([]byte , 0)
	reverse(s , 0 , &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}

}

func reverse(s []byte , i int , res *[]byte){
	if i == len(s) {
		return
	}
	reverse(s , i + 1 , res)

	*res = append(*res , s[i])
}


// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
func gen(n int)[]*TreeNode{
	if n <= 0 {
		return nil
	}
	return genNode(1 ,n)
}
func genNode(start, end int)[]*TreeNode{
	if start > end {
		return []*TreeNode{nil}
	}
	ans:=make([]*TreeNode,0)

	for i:=start;i<=end;i++{
		// 递归生成所有左右子树
		lefts:=genNode(start,i-1)
		rights:=genNode(i+1,end)
		// 拼接左右子树后返回
		for j:=0;j<len(lefts);j++{
			for k:=0;k<len(rights);k++{
				root:=&TreeNode{Val:i}
				root.Left=lefts[j]
				root.Right=rights[k]
				ans=append(ans,root)
			}
		}
	}
	return ans


}

// 递归+备忘录
// fibonacci-number
// 斐波那契数

func fib(n int)int{
	return dfs(n)
}
var m map[int]int  = make(map[int]int)
func dfs(n int) int{
	if n < 2 {
		return n
	}

	if m[n] != 0 {
		return m[n]
	}

	ans := dfs(n - 2) + dfs(n - 1)

	// 缓存结果
	m[n] = ans

	return ans
}



