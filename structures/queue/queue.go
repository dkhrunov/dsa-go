package queue

import "container/list"

type Queue struct {
	list *list.List
}

func New() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) EnQueue(val any) {
	q.list.PushBack(val)
}

func (q *Queue) DeQueue() any {
	front := q.list.Front()
	if front == nil {
		return nil
	}
	val := front.Value
	q.list.Remove(front)
	return val
}

func (q *Queue) First() any {
	return q.list.Front().Value
}

func (q *Queue) Last() any {
	return q.list.Back().Value
}

func (q *Queue) Len() int {
	return q.list.Len()
}
