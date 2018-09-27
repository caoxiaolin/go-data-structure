//双向链表

package linkedlist

import (
	//"errors"
	"fmt"
	"math/rand"
	"time"
)

//结构定义
type DoublyLinkNode struct {
	Data  int64           //数据域
	Prior *DoublyLinkNode //前驱指针
	Next  *DoublyLinkNode //后继指针
}

/*************   下面是链表的基本操作   ***********/
//头插法创建链表，新节点始终在第一个位置，即head->next，返回链表头指针
func CreateHeadForDoubly(n int) *DoublyLinkNode {
	fmt.Println("#头插法创建双向链表：")
	p := &DoublyLinkNode{0, nil, nil}
	for i := 0; i < n; i++ {
		q := &DoublyLinkNode{rand.Int63n(1000), p, p.Next}
		p.Next = q
	}
	p.Data = int64(n)
	return p
}

//尾插法创建链表，新节点始终在链表末尾，返回链表头指针
func CreateTailForDoubly(n int) *DoublyLinkNode {
	fmt.Println("#尾插法创建双向  链表：")
	p := &DoublyLinkNode{0, nil, nil}
	q := p
	for i := 0; i < n; i++ {
		q.Next = &DoublyLinkNode{rand.Int63n(1000), q, nil}
		q = q.Next
	}
	p.Data = int64(n)
	return p
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func DoublyLinkedListExample() {
	newLinkNode := CreateHead(10)
	newLinkNode.Traversal()
	newLinkNode = CreateTail(10)
	newLinkNode.Traversal()
}
