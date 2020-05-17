package vector

type Array struct {
	items    []interface{}
	vector   int
	length   int
	capacity int
}

func NewArray() *Array {
	return NewArrayCustom(10)
}

func NewArrayCustom(vector int) *Array {
	return &Array{
		vector: vector,
	}
}

func (a *Array) Append(item interface{}) {
	if a.length >= a.capacity {
		a.resize()
	}
	a.length++
	a.items[a.length-1] = item
}

func (a *Array) Add(item interface{}, index int) bool {
	if index > a.length {
		return false
	}

	if index == a.length {
		a.Append(item)
		return true
	}

	if a.length >= a.capacity {
		items := a.allocate()
		copy(items[0:index], a.items[0:index])
		items[index] = item
		copy(items[index+1:a.length+1], a.items[index:a.length])
		a.items = items
		a.length++
		a.capacity++

		return true
	}

	copy(a.items[index+1:a.length+1], a.items[index:a.length])
	a.items[index] = item
	a.length++

	return true
}

func (a *Array) Set(item interface{}, index int) bool {
	if index > a.length-1 {
		return false
	}

	a.items[index] = item

	return true
}

func (a *Array) Get(index int) interface{} {
	val, _ := a.GetExisted(index)
	return val
}

func (a *Array) GetExisted(index int) (interface{}, bool) {
	if index > a.length-1 {
		return nil, false
	}

	return a.items[index], true
}

func (a *Array) Len() int {
	return a.length
}

func (a *Array) Cap() int {
	return a.capacity
}

func (a *Array) Remove(index int) (interface{}, bool) {
	if index > a.length-1 {
		return nil, false
	}

	val := a.items[index]

	for i := index; i < a.length-1; i++ {
		tmp := a.items[i]
		a.items[i] = a.items[i+1]
		a.items[i+1] = tmp
	}

	a.length--

	// free memory
	if a.length+2*a.vector <= a.capacity {
		items := a.allocate()
		copy(items[0:a.length], a.items[0:a.length])
		a.items = items
		a.capacity = a.length
	}

	return val, true
}

func (a *Array) allocate() []interface{} {
	return make([]interface{}, a.length+a.vector)
}

func (a *Array) resize() {
	items := a.allocate()
	copy(items, a.items)
	a.items = items
	a.capacity = len(items)
}
