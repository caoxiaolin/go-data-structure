//队列，顺序存储

package queue

import (
	"errors"
	"fmt"
)

//队列的长度
const SIZE = 10

type Queue struct {
	Data [SIZE]interface{}
	Rear int
}

//初始化队列，空队列Rear=-1
func InitQueue() *Queue {
	s := new(Queue)
	s.Rear = -1
	return s
}

//入队，O（1）
func (s *Queue) Enter(e interface{}) error {
	if s.Rear == SIZE-1 {
		return errors.New("Queue is full")
	}
	s.Rear++
	s.Data[s.Rear] = e
	//fmt.Printf("%v 入队\n", e)
	return nil
}

//出队，O（n）
func (s *Queue) Out() (interface{}, error) {
	if s.Rear == -1 {
		return nil, errors.New("Queue is empty")
	}
	r := s.Data[0] //队首元素
	//后序元素依次向前移动1位
	for i := 0; i < s.Rear; i++ {
		s.Data[i] = s.Data[i+1]
	}
	s.Data[s.Rear] = nil
	s.Rear--
	//fmt.Printf("%v 出队\n", r)
	return r, nil
}

//是否空队列
func (s *Queue) IsEmpty() bool {
	if s.Rear == -1 {
		return true
	}
	return false
}

//输出队列
func (s *Queue) Print() {
	for i := 0; i <= s.Rear; i++ {
		fmt.Printf("%v\t", s.Data[i])
	}
	fmt.Println()
}

func Example() {
	queue := InitQueue()
	for i := 1; i <= 10; i++ {
		queue.Enter(i)
	}
	fmt.Println("初始化队列：")
	queue.Print()
	err1 := queue.Enter(11)
	if err1 != nil {
		fmt.Println(err1)
	}
	queue.Print()
	data2, err2 := queue.Out()
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Out element %v\n", data2)
	}
	data3, err3 := queue.Out()
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Printf("Out element %v\n", data3)
	}
	queue.Print()
	fmt.Printf("%+v\n", queue)
}
