package queue

import "errors"

var (
	ErrorEmptyQueue = errors.New("queue is empty")
)

type Queue struct {
	queue []int
}

func Constructor() Queue {
	return Queue{
		queue: make([]int, 0),
	}
}

func (qu *Queue) Push(val int) {
	qu.queue = append(qu.queue, val)
}

func (qu *Queue) Pop() error {
	if len(qu.queue) == 0 {
		return ErrorEmptyQueue
	}
	qu.queue = qu.queue[:0]
	return nil
}

func (qu *Queue) Top() (int, error) {
	if len(qu.queue) == 0 {
		return 0, ErrorEmptyQueue
	}
	return qu.queue[0], nil
}
