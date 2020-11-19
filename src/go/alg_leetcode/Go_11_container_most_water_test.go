package alg_leetcode

import "math"

/**
 * @author
 * @date 2020/11/19 16:04
 * create by Goland
 * @version 1.0
 */ 

// 每个数组代表一个高度，选两个任意的柱子往里边倒水，能最多倒多少水。

func maxArea(arr []int)int{
	max , l , r := 0 , 0 ,len(arr) -1

	for l < r {
		max = int( math.Max(float64(max), math.Min(float64(arr[l]), float64(arr[r]))*float64(r-l)) )

		if arr[l] < arr[r] {
			l++
		}else {
			r--
		}
	}

	return max
}