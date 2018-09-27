//单链表

package linkedlist

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/**
         ┌────────────┐    ┌───────────┐    ┌───────────┐    ┌───────────┐
NULL <---| | HEAD | •-|----|   | A | •-|----|-> | B | •-|----|-> | C | •-|---> NULL
         └────────────┘    └───────────┘    └───────────┘    └───────────┘
*/

//结构定义
type LinkNode struct {
	Data int64     //数据域
	Next *LinkNode //指针域
}

/*************   下面是链表的基本操作   ***********/
//头插法创建链表，新节点始终在第一个位置，即head->next，返回链表头指针
func CreateHead(n int) *LinkNode {
	fmt.Println("#头插法创建链表：")
	p := &LinkNode{0, nil}
	for i := 0; i < n; i++ {
		q := &LinkNode{rand.Int63n(1000), p.Next}
		p.Next = q
	}
	p.Data = int64(n)
	return p
}

//尾插法创建链表，新节点始终在链表末尾，返回链表头指针
func CreateTail(n int) *LinkNode {
	fmt.Println("#尾插法创建链表：")
	p := &LinkNode{0, nil}
	q := p
	for i := 0; i < n; i++ {
		q.Next = &LinkNode{rand.Int63n(1000), nil}
		q = q.Next
	}
	p.Data = int64(n)
	return p
}

//插入数据，指定位置和data，查找位置O(n)，插入操作O(1)
func (head *LinkNode) Insert(i int, data int64) error {
	fmt.Printf("#在第 %d 个位置插入 %d：\n", i, data)
	if i < 1 || i > int(head.Data)+1 {
		return errors.New(fmt.Sprintf("The location to insert does not exist, location is %d\n", i))
	}
	p := head
	j := 1
	for p.Next != nil && j < i {
		p = p.Next
		j++
	}
	if p.Next == nil && j > i {
		return errors.New(fmt.Sprintf("The location to insert does not exist, location is %d\n", i))
	}
	q := &LinkNode{data, p.Next}
	p.Next = q
	head.Data++
	return nil
}

//获取指定的元素，O(n)
func (head *LinkNode) GetElem(i int) (int64, error) {
	fmt.Printf("#获取第 %d 个元素：\n", i)
	err := errors.New("Element is not exists")
	if i < 1 || i > int(head.Data) {
		return 0, err
	}
	p := head.Next
	j := 1
	for p != nil && j <= i {
		if j == i {
			return p.Data, nil
		}
		j++
		p = p.Next
	}
	return 0, err
}

//删除数据，指定位置，查找位置O(n)，删除操作O(1)
func (head *LinkNode) Delete(i int) (int64, error) {
	fmt.Printf("#删除第 %d 个元素：\n", i)
	if i < 1 || i > int(head.Data) {
		return 0, errors.New(fmt.Sprintf("Element %d is not exists\n", i))
	}
	p := head
	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}
	q := p.Next
	p.Next = q.Next
	head.Data--
	return q.Data, nil
}

//链表反转，O(n)
func (head *LinkNode) Reverse() bool {
	fmt.Println("#链表反转：")
	p := head.Next
	q, pn := p, p
	for p != nil {
		pn = p.Next
		if p == head.Next {
			p.Next = nil
		} else {
			p.Next = q
		}
		q = p
		p = pn
	}
	head.Next = q
	return true
}

//数据交换，指定两个位置，查找位置O(n)，交换操作O(1)
func (head *LinkNode) Swap(i, j int) bool {
	fmt.Printf("#交换第 %d 位和 %d 位的元素：\n", i, j)
	if i == j || i > int(head.Data) || j > int(head.Data) {
		return false
	}
	p, q1, q2, q3 := head, head, head, head
	n := 0
	for n < i || n < j {
		if n == i-1 {
			q1 = p
			q2 = p.Next.Next
		}
		if n == j-1 {
			q3 = p
		}
		p = p.Next
		n++
	}
	if q1.Next == q3 {
		q1.Next = q2
		q3.Next = q2.Next
		q2.Next = q3
	} else {
		q1.Next.Next = p.Next
		q3.Next = q1.Next
		q1.Next = p
		p.Next = q2
	}
	return true
}

//遍历链表
func (head *LinkNode) Traversal() {
	p := head.Next
	for p != nil {
		fmt.Printf("%d\t", p.Data)
		p = p.Next
	}
	fmt.Println()
}

/***********   链表基本操作结束   **********/

/***********   下面是链表的排序   **********/
//冒泡排序
func (head *LinkNode) BubbleSort() bool {
	fmt.Println("#冒泡排序：")
	p := head
	var q *LinkNode
	q = nil
	if p.Next == nil || p.Next.Next == nil {
		return false
	}
	for p.Next != q {
		for p.Next.Next != q {
			if p.Next.Data > p.Next.Next.Data {
				x := p.Next
				y := p.Next.Next
				p.Next = y
				x.Next = y.Next
				y.Next = x
			}
			p = p.Next
		}
		q = p.Next
		p = head
	}
	return true
}

/***********   链表排序操作结束   **********/

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Example() {
	newLinkNode := CreateHead(10)
	newLinkNode.Traversal()
	newLinkNode = CreateTail(10)
	newLinkNode.Traversal()
	err1 := newLinkNode.Insert(3, 100)
	if err1 != nil {
		fmt.Print(err1)
	}
	newLinkNode.Traversal()

	elem2, err2 := newLinkNode.GetElem(7)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("get %d\n", elem2)
	}

	elem3, err3 := newLinkNode.Delete(1)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Printf("delete %d\n", elem3)
	}
	newLinkNode.Traversal()

	newLinkNode.Reverse()
	newLinkNode.Traversal()

	newLinkNode.Swap(2, 6)
	newLinkNode.Traversal()

	newLinkNode.BubbleSort()
	newLinkNode.Traversal()
}
