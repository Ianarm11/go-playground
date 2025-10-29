package Stack

import "errors"

var (
	ErrorEmptyStack = errors.New("stack is empty")
)

type StringStack struct {
	stack []string
}

func Constructor() StringStack {
	return StringStack{
		stack: make([]string, 0),
	}
}

func (ms *StringStack) Push(val string) {
	ms.stack = append(ms.stack, val)
}

func (ms *StringStack) Pop() error {
	if len(ms.stack) == 0 {
		return ErrorEmptyStack
	}
	ms.stack = ms.stack[:len(ms.stack)-1]

	return nil
}

func (ms *StringStack) Top() (string, error) {
	if len(ms.stack) == 0 {
		return "", ErrorEmptyStack
	}
	return ms.stack[len(ms.stack)-1], nil
}

type IntStack struct {
	stack []int
}

func Constructor_Int() IntStack {
	return IntStack{
		stack: make([]int, 0),
	}
}

func (ms *IntStack) Push(val int) {
	ms.stack = append(ms.stack, val)
}

func (ms *IntStack) Pop() error {
	if len(ms.stack) == 0 {
		return ErrorEmptyStack
	}
	ms.stack = ms.stack[:len(ms.stack)-1]

	return nil
}

func (ms *IntStack) Top() (int, error) {
	if len(ms.stack) == 0 {
		return -5000, ErrorEmptyStack
	}
	return ms.stack[len(ms.stack)-1], nil
}
