//栈，顺序存储

package stack

import (
	"errors"
	"fmt"
)

//栈的长度
const SIZE = 10

type Stack struct {
	Data [SIZE]interface{}
	Top  int
}

//初始化空栈，空栈Top=-1
func InitStack() *Stack {
	s := new(Stack)
	s.Top = -1
	return s
}

//入栈
func (s *Stack) Push(e interface{}) error {
	if s.Top == SIZE-1 {
		return errors.New("stack is full")
	}
	s.Top++
	s.Data[s.Top] = e
	//fmt.Printf("%v 入栈\n", e)
	return nil
}

//出栈
func (s *Stack) Pop() (interface{}, error) {
	if s.Top == -1 {
		return nil, errors.New("stack is empty")
	}
	r := s.Data[s.Top]
	s.Data[s.Top] = nil
	s.Top--
	//fmt.Printf("%v 出栈\n", r)
	return r, nil
}

//获取栈顶元素
func (s *Stack) GetTop() (interface{}, error) {
	if s.Top == -1 {
		return nil, errors.New("stack is empty")
	}
	return s.Data[s.Top], nil
}

//是否空栈
func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

//输出栈内元素
func (s *Stack) Print() {
	for i := 0; i <= s.Top; i++ {
		fmt.Printf("%v\t", s.Data[i])
	}
	fmt.Println()
}

func Example() {
	stack := InitStack()
	for i := 1; i <= 10; i++ {
		stack.Push(i)
	}
	err1 := stack.Push(11)
	if err1 != nil {
		fmt.Println(err1)
	}
	stack.Print()
	data, err2 := stack.Pop()
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Pop element %v\n", data)
	}
	stack.Print()
}
