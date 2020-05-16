package factor

type Factor struct {
	items    []interface{}
	factor   int
	length   int
	capacity int
}

func NewFactor() *Factor {
	return NewFactorCustom(2)
}

func NewFactorCustom(factor int) *Factor {
	return &Factor{
		factor: factor,
	}
}

func (f *Factor) Append(item interface{}) {
	if f.length >= f.capacity {
		f.resize()
	}

	f.items[f.length-1] = item
}

func (f *Factor) Add(item interface{}, index int) bool {
	if index > f.length {
		return false
	}

	if index == f.length {
		f.Append(item)
		return true
	}

	if f.length >= f.capacity {
		items := f.allocate()
		copy(items[0:index], f.items[0:index])
		items[index] = item
		copy(items[index+1:f.length+1], f.items[index:f.length])
		f.items = items
		f.length++
		f.capacity++

		return true
	}

	copy(f.items[index+1:f.length+1], f.items[index:f.length])
	f.items[index] = item
	f.length++

	return true
}

func (f *Factor) Set(item interface{}, index int) bool {
	if index > f.length-1 {
		return false
	}

	f.items[index] = item

	return true
}

func (f *Factor) Get(index int) interface{} {
	val, _ := f.GetExisted(index)
	return val
}

func (f *Factor) GetExisted(index int) (interface{}, bool) {
	if index > f.length-1 {
		return nil, false
	}

	return f.items[index], true
}

func (f *Factor) Len() int {
	return f.length
}

func (f *Factor) Cap() int {
	return f.capacity
}

func (f *Factor) Remove(index int) (interface{}, bool) {
	if index > f.length-1 {
		return nil, false
	}

	val := f.items[index]

	for i := index; i < f.length-1; i++ {
		tmp := f.items[i]
		f.items[i] = f.items[i+1]
		f.items[i+1] = tmp
	}

	f.length--

	// free memory
	if f.length*2*f.factor <= f.capacity {
		items := f.allocate()
		copy(items[0:f.length], f.items[0:f.length])
		f.items = items
		f.capacity = f.length
	}

	return val, true
}

func (f *Factor) allocate() []interface{} {
	return make([]interface{}, f.factor*f.length+1)
}

func (f *Factor) resize() {
	items := make([]interface{}, f.factor*f.length+1)
	copy(items, f.items)
	f.items = items
	f.capacity++
	f.length++
}
