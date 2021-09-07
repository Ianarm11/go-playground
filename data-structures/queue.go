package data_structures

import "fmt"

type Queue []interface{}

func (Q Queue) Enqueue(data interface{}) []interface{} {
	return append(Q, data)
}

func (Q Queue) Dequeue() []interface{} {
	return Q[1:]
}

func (Q Queue) Display() {
	fmt.Println(Q)
}
