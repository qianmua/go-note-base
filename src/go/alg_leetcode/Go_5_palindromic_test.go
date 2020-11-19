package alg_leetcode

import "math"

/**
 * @author
 * @date 2020/11/19 14:17
 * create by Goland
 * @version 1.0
 */

// 回文字符串

// 给定一个字符串，输出最长的回文子串。回文串指的是正的读和反的读是一样的字符串，例如 "aba"，"ccbbcc"。

// eg:
// 扩展中心
// 每次循环选择一个中心，进行左右扩展，判断左右字符是否相等即可。
func  longest(s string) string{
	if s =="" || len(s) < 1 {
		return ""
	}

	start , end := 0 , 0
	for i := 0; i < len(s);i++  {
		l1 := expandAroundCenter(s , i , i)
		l2 := expandAroundCenter(s , i , i + 1)

		len1 := int(math.Max(float64(l1) , float64(l2)))

		if len1 > end - start {
			start = i  - (len1 - 1) / 2
			end = i + len1 /2
		}
	}

	return s[start: end-1]
}

func expandAroundCenter(s string , left , right int)int{
	L , R := left , right

	for L >= 0 && R < len(s) && s[R] == s[L] {
		L--
		R++
	}
	return R - L - 1
}



 
