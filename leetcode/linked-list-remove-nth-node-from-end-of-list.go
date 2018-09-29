package leetcode

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.

说明：
给定的 n 保证是有效的。

进阶：
你能尝试使用一趟扫描实现吗？

@see https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//构造一个头节点，指向head，这样统一链表操作
	p := &ListNode{Val: 0, Next: head}
	//p、q均指向head
	q := p
	qhead := q
	//p先向后移动n步，此时，p、q间隔为n
	for i := 0; i < n; i++ {
		p = p.Next
	}
	//p和q同时向后移动，直到p到达链表结尾
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	//删除节点，q指向下一个节点的下一个节点
	q.Next = q.Next.Next

	//跳过构造的头节点，返回q的head
	return qhead.Next
}
