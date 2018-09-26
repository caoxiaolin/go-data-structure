//循环队列

package queue

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	Data [SIZE]interface{}
	Head int
	Rear int
}

//初始化队列，空队列 Head = -1, Rear = -1
func InitCircularQueue() *CircularQueue {
	s := new(CircularQueue)
	s.Head = -1
	s.Rear = -1
	return s
}

//入队，O（1）
func (s *CircularQueue) EnterCircularQueue(e interface{}) error {
	if s.Rear-s.Head == SIZE-1 || s.Head-s.Rear == 1 {
		return errors.New("Queue is full")
	}
	if s.Head == -1 {
		s.Head = 0
	}
	if s.Rear == SIZE-1 {
		s.Rear = 0
	} else {
		s.Rear++
	}
	s.Data[s.Rear] = e
	//fmt.Printf("%v 入队\n", e)
	return nil
}

//出队，O（1）
func (s *CircularQueue) OutCircularQueue() (interface{}, error) {
	if s.Rear == -1 {
		return nil, errors.New("Queue is empty")
	}
	r := s.Data[s.Head] //队首元素
	s.Data[s.Head] = nil
	if s.Head == s.Rear { //只有一个元素，队列置空
		s.Head = -1
		s.Rear = -1
	} else if s.Head == SIZE-1 { //head指向末尾，重头开始
		s.Head = 0
	} else {
		s.Head++
	}
	//fmt.Printf("%v 出队\n", r)
	return r, nil
}

//输出队列
func (s *CircularQueue) PrintCircularQueue() {
	if s.Rear == -1 {
		return
	}
	max := s.Rear
	if s.Rear < s.Head {
		max = SIZE - 1
	}
	for i := s.Head; i <= max; i++ {
		fmt.Printf("%v\t", s.Data[i])
	}
	if s.Rear < s.Head {
		for i := 0; i <= s.Rear; i++ {
			fmt.Printf("%v\t", s.Data[i])
		}
	}
	fmt.Println()
}

func Example1() {
	queue := InitCircularQueue()
	for i := 1; i <= 10; i++ {
		queue.EnterCircularQueue(i)
	}
	fmt.Println("初始化队列：")
	queue.PrintCircularQueue()
	fmt.Printf("%+v\n", queue)
	data1, err1 := queue.OutCircularQueue()
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("Out element %v\n", data1)
	}
	queue.PrintCircularQueue()
	fmt.Printf("%+v\n", queue)
	data2, err2 := queue.OutCircularQueue()
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Out element %v\n", data2)
	}
	queue.PrintCircularQueue()
	fmt.Printf("%+v\n", queue)
	err3 := queue.EnterCircularQueue(11)
	if err3 != nil {
		fmt.Println(err1)
	}
	queue.PrintCircularQueue()
	fmt.Printf("%+v\n", queue)
}
