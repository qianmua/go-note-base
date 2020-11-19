package alg_leetcode

import "sort"

/**
 * @author
 * @date 2020/11/19 17:13
 * create by Goland
 * @version 1.0
 */

// 三数和为0

//  0 减去当前的数
// 两个指针

// 保证不加入重复的 list

func threeSum(num []int) (res [][]int) {

	sort.Ints(num)

	res = make([][]int ,0)
	for i := 0; i < len(num) -2;i++ {
		if i == 0 || (i > 0 && num[i] != num[i-1]) {
			//两个指针,并且头指针从i + 1开始，防止加入重复的元素
			lo , hi := i+1 , len(num) -1
			sum := 0 - num[i]

			for lo < hi {
				if num[lo] + num[hi] == sum {

					res[i] = append(res[i] , num[i] , num[lo] , num[hi])

					// 元素相同要后移，防止加入重复的 list
					for lo < hi && num[lo] == num[lo + 1] {
						lo ++ 						
					}
					for lo < hi && num[hi] == num[hi - 1] {
						hi --
					}
					lo ++
					hi --

				}else if num[lo] + num[hi] < sum {
					lo ++
				}else {
					hi --
				}

			}
		}
	}
	return
}

func test1(args ... int){

}

func test2(){
	test1(1,2,3,4,5)
}
 
