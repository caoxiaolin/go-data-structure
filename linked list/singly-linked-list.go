//单链表

package main

import (
    "fmt"
    "math/rand"
    "time"
)

//结构定义
type LinkNode struct{
    Data int64
    Next *LinkNode
}

//初始化，生成头指针，头指针data存储链表长度
func New() *LinkNode{
    return &LinkNode{0, nil}
}

/*************   下面是链表的基本操作   ***********/
//插入数据，指定位置和data
func (head *LinkNode) Insert(i int, data int64) bool{
    p := head
    j := 1
    for p.Next != nil && j < i {
        p = p.Next
        j++
    }
    n := new(LinkNode)
    n.Data = data
    n.Next = p.Next
    p.Next = n
    head.Data++
    return true
}

//获取指定的元素
func (head *LinkNode) GetElem(i int) int64{
    p := head
    for p.Next != nil{
        if i == 0 {
              return p.Data
        }
        i--
        p = p.Next
    }
    return -1
}

//删除数据，指定位置
func (head *LinkNode) Delete(i int) bool{
    if i < 1 || i > int(head.Data){
        return false
    }
    p := head
    j := 1
    for p.Next != nil && j < i {
        p = p.Next
        j++
    }
    p.Next = p.Next.Next
    head.Data--
    return true
}

//数据交换，指定两个位置
func (head *LinkNode) Swap(i, j int) bool{
    if i == j || i > int(head.Data) || j > int(head.Data){
        return false
    }
    p, q1, q2, q3 := head, head, head, head
    n := 0
    for n < i || n < j{
        if n == i - 1{
            q1 = p
            q2 = p.Next.Next
        }
        if n == j -1 {
            q3 = p
        }
        p = p.Next
        n++
    }
    if q1.Next == q3{
        q1.Next = q2
        q3.Next = q2.Next
        q2.Next = q3
    }else{
        q1.Next.Next = p.Next
        q3.Next = q1.Next
        q1.Next = p
        p.Next = q2
    }
    return true
}

//数据交换，指定两个位置
func (head *LinkNode) Swap1(i, j int) bool{
    if i == j || i > int(head.Data) || j > int(head.Data){
        return false
    }
    p := head
    pi, pj := p, p
    n := 1
    for n <= i || n <= j{
        p = p.Next
        if n == i{
            pi = p
        }
        if n == j{
            pj = p
        }
        n++
    }
    head.Delete(i)
    head.Insert(i, pj.Data)
    head.Delete(j)
    head.Insert(j, pi.Data)
    return true
}

//链表反转
func (head *LinkNode) Reverse() bool{
    p := head.Next
    q, pn := p, p
    for p != nil{
        pn = p.Next
        if p == head.Next{
            p.Next = nil
        }else{
            p.Next = q
        }
        q = p
        p = pn
    }
    head.Next = q
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
func (head *LinkNode) Sort(){

}

/***********   链表排序操作结束   **********/

func init(){
    rand.Seed(time.Now().UnixNano())
}


func main(){
    newLinkNode := New()
    for n := 1; n <= 3; n++ {
        newLinkNode.Insert(1, rand.Int63n(1000))
    }

    newLinkNode.Traversal()

    newLinkNode.Swap(1, 3)
    newLinkNode.Traversal()

    newLinkNode.Delete(5)
    newLinkNode.Traversal()

    newLinkNode.Reverse()
    newLinkNode.Traversal()

    fmt.Printf("%d\n", newLinkNode.GetElem(2))
}
