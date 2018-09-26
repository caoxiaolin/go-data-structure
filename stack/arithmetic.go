//栈的应用，加减乘除四则运算，逆波兰
package stack

import (
	"fmt"
	"strconv"
)

type Arithmetic struct {
	Input            string   //原始输入
	MiddleExpression []string //中序表达式
	SuffixExpression []string //后序表达式
}

//规范表达式，处理多位数，输出到中序表达式
func (a *Arithmetic) Format() {
	if a.Input != "" {
		strlen := len(a.Input)
		var str string
		for i := 0; i < strlen; i++ {
			c := a.Input[i]
			s := string(c)
			//如果是数字，则继续往后判断
			if a.IsNumber(c) {
				str += s //多位数连接
				if i+1 < strlen && a.IsNumber(a.Input[i+1]) {
					continue
				}
			} else {
				if str != "" {
					a.MiddleExpression = append(a.MiddleExpression, str)
				}
				a.MiddleExpression = append(a.MiddleExpression, s)
				str = ""
			}
			if i+1 == strlen && str != "" {
				a.MiddleExpression = append(a.MiddleExpression, str)
			}
		}
	}
}

//转为逆波兰表达式（后缀表达式）
//规则：顺序遍历，数字输出，符号则判断与栈顶符号的优先级，右括号或优先级不高于栈顶符号
//栈顶元素依次出栈，遇左括号停止，当前符号进栈，最后都出栈
func (a *Arithmetic) RPN() {
	if a.Input != "" {
		a.Format()
		stack := InitStack()
		for _, s := range a.MiddleExpression {
			//fmt.Printf("当前 s = %v\n", s)
			//转换为int，如果转换失败，说明是括号或操作符
			_, err := strconv.Atoi(s)
			if err != nil { //转换失败
				top, _ := stack.GetTop()
				//s为"("时，直接入栈
				if s == "(" {
					stack.Push(s)
					continue
				}
				//s为右括号，或者s运算符的优先级不高于栈顶运算符优先级
				for top != nil && (s == ")" || top.(string) == "*" || top.(string) == "/") {
					//栈顶元素开始出栈，直到"("出栈为止
					e, _ := stack.Pop()
					v, _ := e.(string)
					if v == "(" {
						break
					}
					a.SuffixExpression = append(a.SuffixExpression, v)
					top, _ = stack.GetTop()
				}
				//当前符号进栈
				if s != ")" {
					stack.Push(s)
				}
			} else {
				//数字直接入栈
				a.SuffixExpression = append(a.SuffixExpression, s)
			}
		}

		//如果栈不为空，则元素全部出栈
		for stack.IsEmpty() != true {
			e, _ := stack.Pop()
			a.SuffixExpression = append(a.SuffixExpression, e.(string))
		}
	}
}

//计算最终结果
//数字挨个入栈，遇到符号，前两个数出栈并计算，结果入栈，直到计算出最终结果出栈
func (a *Arithmetic) CalcResult() float64 {
	if a.Input != "" {
		a.RPN()
		stack := InitStack()
		for _, s := range a.SuffixExpression {
			val, err := strconv.ParseFloat(s, 64)
			if err == nil {
				stack.Push(val)
			} else {
				n, _ := stack.Pop()
				m, _ := stack.Pop()
				stack.Push(a.Calc(m.(float64), n.(float64), s))
			}
		}
		res, _ := stack.Pop()
		return res.(float64)
	}
	return 0
}

//判断是否数字
func (a *Arithmetic) IsNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}

//基本的加减乘除四则运算
func (a *Arithmetic) Calc(x float64, y float64, o string) (res float64) {
	switch o {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		res = x / y
	}
	//fmt.Printf("%v %v %v = %v\n", x, o, y, res)
	return
}

func Arith() {
	a := new(Arithmetic)
	a.Input = "11+2*(42-3)/5*((7-6)/8*9)"
	fmt.Printf("%s = %f\n", a.MiddleExpression, a.CalcResult())
}
