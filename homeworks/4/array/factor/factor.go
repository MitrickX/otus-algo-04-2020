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

func (f *Factor) Add(item interface{}) {
	if f.length >= f.capacity {
		f.resize()
	}

	f.items[f.length-1] = item
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

	return val, true
}

func (f *Factor) resize() {
	items := make([]interface{}, f.factor*f.length+1)
	copy(items, f.items)
	f.items = items
	f.capacity++
	f.length++
}
