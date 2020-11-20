package alg_leetcode

/**
 * @author
 * @date 2020/11/20 17:27
 * create by Goland
 * @version 1.0
 */

// 给定一个链表，将倒数第 n 个结点删除。

// 双指针
func removeNodeEnd1(head *ListNode , n int) *ListNode{
	var dummy *ListNode

	dummy.Next = head
	first := dummy
	second := dummy

	// first
	for i := 1; i<= n + 1 ;i ++  {
		first = first.Next
	}

	for first != nil{
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next

}
