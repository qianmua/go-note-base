package alg_leetcode

/**
 * @author
 * @date 2020/11/19 16:46
 * create by Goland
 * @version 1.0
 */ 
 
// 最长相同前缀
func longPrefix(strs []string)string{
	if strs == nil || len(strs) == 0 {
		return ""
	}

	// each
	for i:= 0 ; i < len(strs[0]) ; i ++{
		var c byte = strs[0][i]
		// 比较后面

		for j := 1 ; j < len(strs) ; j ++{

			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}