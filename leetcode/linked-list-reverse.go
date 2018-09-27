package leetcode

/**
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。
示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
@see https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
*/

import (
	"github.com/caoxiaolin/go-data-structure/linked-list"
)

func ReverseBetween(head *linkedlist.LinkNode, m int, n int) *linkedlist.LinkNode {
	if m == n {
		return head
	}
	prev := head
	cur := head
	next := head
	start := head
	end := head
	i := 0
	for i < m {
		//i==m-1时，表示从下一个位置开始要反转，所以记录当前p的位置，在反转完成后修改x.Next
		if i == m-1 {
			start = cur
			end = cur.Next
		}
		cur = cur.Next
		i++
	}
	//开始反转
	for i <= n {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		i++
	}
	start.Next = prev
	end.Next = cur
	return head
}

func TestReverseBetween() {
	list := linkedlist.CreateHead(10)
	list.Traversal()
	ReverseBetween(list, 6, 10)
	list.Traversal()
}
