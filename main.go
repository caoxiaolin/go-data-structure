package main

import (
	"fmt"
	"github.com/caoxiaolin/go-data-structure/linked-list"
	"github.com/caoxiaolin/go-data-structure/queue"
	"github.com/caoxiaolin/go-data-structure/stack"
	"github.com/caoxiaolin/go-data-structure/tree"
)

func main() {
	fmt.Println("####################  链表  #####################")
	linkedlist.Example()
	linkedlist.DoublyLinkedListExample()
	fmt.Println("####################  队列  #####################")
	queue.Example()
	queue.Example1()
	fmt.Println("####################  堆栈  #####################")
	stack.Example()
	stack.Arith()
	fmt.Println("####################  树  #####################")
	btree.Example()
}
