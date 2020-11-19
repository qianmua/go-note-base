package alg_leetcode

import "math"

/**
 * @author
 * @date 2020/11/19 15:10
 * create by Goland
 * @version 1.0
 */ 


// 倒置数字

// 输入整数，输出它的倒置。

func reverseNumber(n int)int32{
	rev := int64(n)

	for n != 0{
		pop := n % 10
		n /=10
		rev = rev * 10 + int64(pop)
	}

	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	return int32(rev)
}
