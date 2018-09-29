package leetcode

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

@see https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	res := new(ListNode)
	tmp := res
	for p != nil && q != nil {
		if p.Val < q.Val {
			tmp.Next = p
			p = p.Next
			tmp = tmp.Next
		} else {
			tmp.Next = q
			q = q.Next
			tmp = tmp.Next
		}
	}
	if p != nil {
		tmp.Next = p
	}
	if q != nil {
		tmp.Next = q
	}
	return res.Next
}

func TestMergeTowLists() {
	fmt.Println("========== 合并两个有序链表 ============")
	//构造链表
	list1 := []int{18, 20, 25, 27, 32, 48}
	nodehead1 := &ListNode{Val: 0, Next: nil}
	listnode1 := nodehead1
	fmt.Print("链表1：")
	for _, v := range list1 {
		listnode1.Next = &ListNode{Val: v, Next: nil}
		listnode1 = listnode1.Next
		fmt.Printf("%d\t", v)
	}
	listnode1 = nodehead1.Next

	list2 := []int{13, 20, 45, 47, 52, 88}
	nodehead2 := &ListNode{Val: 0, Next: nil}
	listnode2 := nodehead2
	fmt.Print("\n链表2：")
	for _, v := range list2 {
		listnode2.Next = &ListNode{Val: v, Next: nil}
		listnode2 = listnode2.Next
		fmt.Printf("%d\t", v)
	}
	listnode2 = nodehead2.Next

	res := mergeTwoLists(listnode1, listnode2)
	//输出结果
	fmt.Println()
	for res != nil {
		fmt.Printf("%d\t", res.Val)
		res = res.Next
	}
	fmt.Println("\n=================================")
}
