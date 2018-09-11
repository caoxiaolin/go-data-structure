package main

import (
    "fmt"
    "math"
    "math/rand"
)

type LinkNode struct{
    Data float64
    Next *LinkNode
}

func New() *LinkNode{
    return &LinkNode{0, nil}
}

func (head *LinkNode) Insert(i int, data float64) bool{
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
    return true
}

func (head *LinkNode) Delete(i int) bool{
    p := head
    j := 1
    for p.Next != nil && j < i {
        p = p.Next
        j++
    }
    if p.Next == nil {
        return false
    }
    p.Next = p.Next.Next
    return true
}

func (head *LinkNode) Traversal() {
    p := head.Next
    for p != nil {
        fmt.Printf("%.f\t", p.Data)
        p = p.Next
    }
    fmt.Println()
}

func main(){
    newLinkNode := New()
    for n := 1; n < 10; n++ {
        newLinkNode.Insert(1, math.Ceil(rand.Float64()*1000))
    }
    newLinkNode.Traversal()

    newLinkNode.Delete(1)
    newLinkNode.Traversal()
}
