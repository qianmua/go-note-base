package alg

/**
 * @author
 * @date 2020/11/17 20:07
 * create by Goland
 * @version 1.0
 */

// 回溯


// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

func subSets(nums []int)[][]int{
	res := make([][]int , 0)

	list := make([]int ,0)

	backtrack(nums , 0 , list , &res)

	return res
}

// pos 下次添加到集合中的元素位置索引
// list 临时结果集合(每次需要复制保存)
func backtrack(nums []int , pos int , list []int , res *[][]int){
	ans := make([]int , len(list))

	copy(ans , list)

	*res = append(*res , ans)

	for i := pos; i <len(nums) ;i ++  {
		list = append(list, nums[i])
		backtrack(nums , i + 1 , list , res)

		list = list[0 : len(list) - 1]
		
	}
}



/// 全排列
func permute(nums []int) [][]int {

	result := make([][]int, 0)
	list := make([]int, 0)

	// 标记这个元素是否已经添加到结果集
	visited := make([]bool, len(nums))
	backtrack2(nums, visited, list, &result)

	return result
}

func backtrack2(nums []int, visited []bool, list []int, result *[][]int) {
	// 返回条件：临时结果和输入集合长度一致 才是全排列
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 已经添加过的元素，直接跳过
		if visited[i] {
			continue
		}

		// 上一个元素和当前相同，并且没有访问过就跳过
		if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}

		// 添加元素
		list = append(list, nums[i])
		visited[i] = true
		backtrack2(nums, visited, list, result)
		// 移除元素
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}


