package queue

import (
	"reflect"
	"testing"
)

func TestPriorityQueue_Enqueue(t *testing.T) {
	q := NewPriorityQueue()

	if q.Len() != 0 {
		t.Fatalf("unepxected that length %d isntread of %d", q.Len(), 0)
	}

	q.Enqueue(1, 1)
	q.Enqueue(0, 2)
	q.Enqueue(2, 3)

	if q.Len() != 3 {
		t.Fatalf("unepxected that length %d isntread of %d", q.Len(), 3)
	}
}

func TestPriorityQueue_Dequeue(t *testing.T) {
	q := NewPriorityQueue()

	_, ok := q.DequeueExisted()

	if ok {
		t.Fatalf("unepxected ok dequeue from empty pq")
	}

	q.Enqueue(1, 1)

	val, ok := q.DequeueExisted()

	if !ok {
		t.Fatalf("unepxected not ok dequeue from not epmty pq")
	}

	if val != 1 {
		t.Fatalf("unepxected dequeue-ed value %d instread of %d", val, 1)
	}

	if q.Len() != 0 {
		t.Fatalf("unepxected pq is not empty")
	}

	_, ok = q.DequeueExisted()

	if ok {
		t.Fatalf("unepxected ok dequeue from empty pq")
	}
}

func TestPriorityQueue_Discipline(t *testing.T) {
	q := NewPriorityQueue()

	q.Enqueue(3, 1)
	q.Enqueue(1, 2)
	q.Enqueue(2, 3)
	q.Enqueue(0, 4)
	q.Enqueue(4, 5)
	q.Enqueue(0, 6)
	q.Enqueue(0, 7)
	q.Enqueue(2, 8)
	q.Enqueue(10, 9)
	q.Enqueue(1, 10)

	pq := []interface{}(nil)

	for q.Len() != 0 {
		pq = append(pq, q.Dequeue())
	}

	expected := []interface{}{9, 5, 1, 3, 8, 2, 10, 4, 6, 7}

	if !reflect.DeepEqual(pq, expected) {
		t.Fatalf("unexpected not pq discipline: %v instread of %v", pq, expected)
	}
}
