package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	if got := New(); got == nil {
		t.Errorf("NewQueue() can't create new queue struct")
	}
}

func TestQueue_EnQueue(t *testing.T) {
	q := New()
	items := [3]int{1, 2, 3}

	q.EnQueue(items[0])
	if last := q.Last(); !reflect.DeepEqual(items[0], last) {
		t.Errorf("Should enqueue value at the end of the queue; Last = %v, got = %v", items[0], last)
	}
	if first := q.First(); !reflect.DeepEqual(items[0], first) {
		t.Errorf("First item in the queue should be a val1; First = %v, got = %v", items[0], first)
	}

	q.EnQueue(items[1])
	if last := q.Last(); !reflect.DeepEqual(items[1], last) {
		t.Errorf("Should enqueue value at the end of the queue; Last = %v, got = %v", items[1], last)
	}
	if first := q.First(); !reflect.DeepEqual(items[0], first) {
		t.Errorf("Shouldnt change the first of the queue; First = %v, got = %v", items[0], first)
	}

	q.EnQueue(items[2])
	if last := q.Last(); !reflect.DeepEqual(items[2], last) {
		t.Errorf("Should enqueue value at the end of the queue; Last = %v, got = %v", items[2], last)
	}
	if first := q.First(); !reflect.DeepEqual(items[0], first) {
		t.Errorf("Shouldnt change the first of the queue; First = %v, got = %v", items[0], first)
	}

	if q.Len() != len(items) {
		t.Errorf("Should be length = %v, got = %v", len(items), q.Len())
	}
}

func TestQueue_DeQueue(t *testing.T) {
	q := New()
	items := make([]int, 3)
	items[0] = 1
	items[1] = 2
	items[2] = 3

	for _, v := range items {
		q.EnQueue(v)
	}

	i := 0
	deQueuedItems := make([]int, 3)
	for q.Len() > 0 {
		deQueuedItems[i] = q.DeQueue().(int)
		i++
	}

	if q.Len() != 0 {
		t.Errorf("Queue should be empty, but got = %v", q.Len())
	}

	if !reflect.DeepEqual(deQueuedItems, items) {
		t.Errorf("Incorrect order of dequeue items %v, %v", deQueuedItems, items)
	}

	if q.DeQueue() != nil {
		t.Error("Should dequeue return nil, when queue empty")
	}
}
