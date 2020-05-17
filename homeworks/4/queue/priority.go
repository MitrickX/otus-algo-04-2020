package queue

import (
	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/container"
	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/container/list"
)

type pqItem struct {
	*Queue
	priority int
}

type PriorityQueue struct {
	array  container.Array
	length int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		array: list.NewArray(),
	}
}

func (p *PriorityQueue) Enqueue(priority int, item interface{}) {
	// Less, Equal
	index, hit := BinarySearch(p.array,
		func(x interface{}) int {
			return x.(pqItem).priority - priority
		},
	)

	var pq pqItem

	if !hit {
		// pqItem not exists yet for this priority
		pq = pqItem{
			Queue:    NewQueue(),
			priority: priority,
		}
		p.array.Add(pq, index)
	} else {
		// pqItem is in array, extract element from array and assert to Queue
		value := p.array.Get(index)
		pq = value.(pqItem)
	}

	pq.Enqueue(item)

	p.length++
}

func (p *PriorityQueue) DequeueExisted() (interface{}, bool) {
	if p.array.Len() == 0 {
		return nil, false
	}

	// Get last pqItem, cause in the end of array the most priority
	index := p.array.Len() - 1
	value := p.array.Get(index)
	pq := value.(pqItem)

	item := pq.Dequeue()

	if pq.Len() == 0 {
		p.array.Remove(index)
	}

	p.length--

	return item, true
}

func (p *PriorityQueue) Dequeue() interface{} {
	val, ok := p.DequeueExisted()
	if !ok {
		return nil
	}

	return val
}

func (p *PriorityQueue) Len() int {
	return p.length
}
