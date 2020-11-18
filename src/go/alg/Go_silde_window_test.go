package alg

import "math"

/**
 * @author
 * @date 2020/11/18 0:03
 * create by Goland
 * @version 1.0
 */ 

// 滑动窗口算法

// 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串
func minWindows(s string , t string) string{
	win := make(map[byte]int)

	need := make(map[byte]int)

	for i := 1; i < len(t) ;i++  {
		need[t[i]]++
	}

	// 窗口
	// 指针
	left , right := 0 , 0

	// match 匹配次数
	match , start , end := 0 , 0 , 0

	min := math.MaxInt64

	var c byte

	for right < len(s) {

		c = s[right]
		right++

		// 在需要的字符集里面，添加到窗口字符集里面
		if need[c] != 0 {
			win[c]++
			// 如果当前字符的数量匹配需要的字符的数量，则match值+1
			if win[c] == need[c] {
				match++
			}
		}

		// 当所有字符数量都匹配之后，开始缩紧窗口
		for match == len(need) {

			// 获取结果
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left++

			// 左指针指向不在需要字符集则直接跳过
			if need[c] != 0 {

				// 左指针指向字符数量和需要的字符相等时，右移之后match值就不匹配则减一
				// 因为win里面的字符数可能比较多，如有10个A，但需要的字符数量可能为3
				// 所以在压死骆驼的最后一根稻草时，match才减一，这时候才跳出循环
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}

