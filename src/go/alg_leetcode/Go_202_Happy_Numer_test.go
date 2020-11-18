package alg_leetcode

/**
 * @author
 * @date 2020/11/18 17:28
 * create by Goland
 * @version 1.0
 */

// 给一个数字，将这个数字的各个位取平方然后相加，得到新的数字再重复这个过程。如果得到了 1 就返回 true，如果得不到 1 就返回 false 。

func isHappy(n int) bool{
	// 快慢指针
	slow , fast := n , n

	for slow != fast{
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return slow == 1
}

func getNext(n int)int{
	next := 0
	for n > 0{
		 t:= n % 10
		 next += t * t
		 n /= 10
	}


	return next
}
 
