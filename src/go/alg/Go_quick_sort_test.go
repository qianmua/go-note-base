package alg

/**
 * @author
 * @date 2020/11/17 20:06
 * create by Goland
 * @version 1.0
 */

// 快排

func quick(arr []int , s , e int) {
	if s > e {
		return
	}

	left , right := s , e
	mid := arr[left]

	for left < right {

		for left < right && arr[right] >= mid {
			right --
		}

		arr[left] = arr[right]

		for left < right && arr[left] < mid {
			left ++
		}

		arr[right] = arr[left]

	}
	arr[left] = mid
	quick(arr , s , left -1)
	quick(arr ,left + 1 , e)

}