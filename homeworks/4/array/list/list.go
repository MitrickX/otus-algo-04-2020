package list

type List struct {
	items []interface{}
}

func NewList() *List {
	return &List{}
}

func (l *List) Add(item interface{}, index int) bool {
	ln := len(l.items)
	if index > ln {
		return false
	}

	l.Append(item)

	if index == ln {
		return true
	}

	for i := ln; i > index; i-- {
		tmp := l.items[i]
		l.items[i] = l.items[i-1]
		l.items[i-1] = tmp
	}

	return true
}

func (l *List) Append(item interface{}) {
	l.items = append(l.items, item)
}

func (l *List) Set(item interface{}, index int) bool {
	if index > len(l.items)-1 {
		return false
	}

	l.items[index] = item

	return true
}

func (l *List) Get(index int) interface{} {
	val, _ := l.GetExisted(index)
	return val
}

func (l *List) GetExisted(index int) (interface{}, bool) {
	if index > len(l.items)-1 {
		return nil, false
	}

	return l.items[index], true
}

func (l *List) Len() int {
	return len(l.items)
}

func (l *List) Cap() int {
	return cap(l.items)
}

func (l *List) Remove(index int) (interface{}, bool) {
	ln := len(l.items)

	if index > ln-1 {
		return nil, false
	}

	val := l.items[index]

	for i := index; i < ln-1; i++ {
		tmp := l.items[i]
		l.items[i] = l.items[i+1]
		l.items[i+1] = tmp
	}

	// truncate slice
	l.items = l.items[0 : ln-1]

	return val, true
}
