package list

type Array struct {
	items []interface{}
}

func NewArray() *Array {
	return &Array{}
}

func (l *Array) Add(item interface{}, index int) bool {
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

func (l *Array) Append(item interface{}) {
	l.items = append(l.items, item)
}

func (l *Array) Set(item interface{}, index int) bool {
	if index > len(l.items)-1 {
		return false
	}

	l.items[index] = item

	return true
}

func (l *Array) Get(index int) interface{} {
	val, _ := l.GetExisted(index)
	return val
}

func (l *Array) GetExisted(index int) (interface{}, bool) {
	if index > len(l.items)-1 {
		return nil, false
	}

	return l.items[index], true
}

func (l *Array) Len() int {
	return len(l.items)
}

func (l *Array) Cap() int {
	return cap(l.items)
}

func (l *Array) Remove(index int) (interface{}, bool) {
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
