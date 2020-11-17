package alg

import (
	"fmt"
	"testing"
)

/**
 * @author
 * @date 2020/11/18 0:10
 * create by Goland
 * @version 1.0
 */


// 二进制操作

// base
/**
a = 0 ^ a = a ^ 0

0 = a ^ a

即
a = a ^ b ^ b

// 交换两个数字
// a = 1 b = 2
a = a ^ b
b = a ^ b
a = a ^ b



// 移除最后一个 1
a=n&(n-1)
// 获取最后一个 1
diff=(n&(n-1))^n

 */

func TestBinaryOpM1(t *testing.T) {
	a , b := 1,2

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a , b)
}

// 给定一个**非空**整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNum(nums []int)int{
	// a ^ a = 0
	res := 0
	for i := 0; i < len(nums); i++  {
		res = res ^ nums[i]

	}
	return res
}

// 给定一个**非空**整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
func singleNum2(nums []int)int{
	var res int
	// 统计每位1的个数
	for i := 0; i < 64 ; i++ {
		sum := 0
		for j := 0; j < len(nums);j++  {
			sum += (nums[j] >> i) & 1
		}
		// 还原10 = 10^0 |

		res ^= (sum % 3) << i
	}

	return res
}

// 给定一个整数数组  `nums`，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
func singleNum3(nums []int) []int{
	// a=a^b
	// b=a^b
	// a=a^b
	// 把a^b分成两部分,方案：可以通过diff最后一个1区分
	diff := 0
	for i:= 0; i < len(nums) ;i++  {
		diff ^= nums[i]
	}

	result:=[]int{diff,diff}

	// 去掉末尾的1后异或diff就得到最后一个1的位置
	diff=(diff&(diff-1))^diff
	for i:=0;i<len(nums);i++{
		if diff&nums[i]==0{
			result[0]^=nums[i]
		}else{
			result[1]^=nums[i]
		}
	}
	return result
}

// 编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’  的个数
//[汉明重量]
func hanmingWight(n uint32)int{
	res := 0

	for n != 0 {
		n= n & (n - 1)
		res ++
	}

	return res
}
