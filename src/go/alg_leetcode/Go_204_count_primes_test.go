package alg_leetcode

/**
 * @author
 * @date 2020/11/18 17:41
 * create by Goland
 * @version 1.0
 */

// 求出小于 n 的素数个数。

func countPrimes(n int)int{
	iscP := make([]bool , n)
	count := 0

	for i := 2; i< n ;i++  {
		if !iscP[i] {
			count++

			// 将当前素数的倍数依次标记为非素数
			for j := 2 ; j * i < n ; j ++{
				iscP[j * i] = true
			}
		}
	}


	return count
}
 
