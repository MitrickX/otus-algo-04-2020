package queue

type node struct {
	item interface{}
	next *node
}

type Queue struct {
	head   *node
	tail   *node
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item interface{}) {
	n := &node{item: item}

	q.length++

	if q.head == nil {
		q.head = n
		q.tail = q.head

		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *Queue) DequeueExisted() (interface{}, bool) {
	if q.head == nil {
		return nil, false
	}

	n := q.head
	q.head = q.head.next

	n.next = nil
	q.length--

	return n.item, true
}

func (q *Queue) Dequeue() interface{} {
	val, ok := q.DequeueExisted()
	if !ok {
		return nil
	}

	return val
}

func (q *Queue) Len() int {
	return q.length
}
