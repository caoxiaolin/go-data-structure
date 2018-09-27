package leetcode

/**
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
说明：不允许修改给定的链表。
@see https://leetcode-cn.com/problems/linked-list-cycle-ii/description/

基本思路：
快慢指针
*/

import (
	"fmt"
	"github.com/caoxiaolin/go-data-structure/linked-list"
)

func DetectCycle(head *linkedlist.LinkNode) bool {
	//p一次走一步，q一次走两步
	p, q := head, head
	//只有头指针指向自己，有环
	if p.Next == p {
		return true
	}
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
		if p == q {
			return true
		}
	}
	return false
}

func TestCycle() {
	list := linkedlist.CreateHead(10)
	list.Traversal()
	fmt.Printf("%+v\n", DetectCycle(list))
	list.Next.Next.Next.Next.Next.Next.Next = list.Next.Next
	fmt.Printf("%+v\n", DetectCycle(list))
}
