package queue

import (
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue()

	if q.Len() != 0 {
		t.Fatalf("unepxected that length %d isntread of %d", q.Len(), 0)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Len() != 3 {
		t.Fatalf("unepxected that length %d isntread of %d", q.Len(), 3)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue()

	_, ok := q.DequeueExisted()

	if ok {
		t.Fatalf("unepxected ok dequeue from empty pqItem")
	}

	q.Enqueue(1)

	val, ok := q.DequeueExisted()

	if !ok {
		t.Fatalf("unepxected not ok dequeue from not epmty pqItem")
	}

	if val != 1 {
		t.Fatalf("unepxected dequeue-ed value %d instread of %d", val, 1)
	}

	if q.Len() != 0 {
		t.Fatalf("unepxected pqItem is not empty")
	}
}

func TestQueueDiscipline(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	fifo := []interface{}(nil)

	for q.Len() != 0 {
		fifo = append(fifo, q.Dequeue())
	}

	expected := []interface{}{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(fifo, expected) {
		t.Fatalf("unexpected not FIFO discipline: %v instread of %v", fifo, expected)
	}
}
