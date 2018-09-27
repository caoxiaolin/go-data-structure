package leetcode

/**
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。
示例:
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
@see https://leetcode-cn.com/problems/partition-list/description/
*/

import (
	"github.com/caoxiaolin/go-data-structure/linked-list"
)

func PartitionList(head *linkedlist.LinkNode, n int) *linkedlist.LinkNode {
	x := &linkedlist.LinkNode{Data: 0, Next: nil}
	y := &linkedlist.LinkNode{Data: 0, Next: nil}
	xh, yh := x, y
	p := head
	for p.Next != nil {
		if p.Next.Data < int64(n) {
			x.Next = &linkedlist.LinkNode{Data: p.Next.Data, Next: nil}
			x = x.Next
		} else {
			y.Next = &linkedlist.LinkNode{Data: p.Next.Data, Next: nil}
			y = y.Next
		}
		p = p.Next
	}
	x.Next = yh.Next
	return xh
}

func TestPartition() {
	list := linkedlist.CreateHead(10)
	list.Traversal()
	list1 := PartitionList(list, 700)
	list1.Traversal()
}
