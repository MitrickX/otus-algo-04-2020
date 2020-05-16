package vector

type Vector struct {
	items    []interface{}
	vector   int
	length   int
	capacity int
}

func NewVector() *Vector {
	return NewVectorCustom(10)
}

func NewVectorCustom(vector int) *Vector {
	return &Vector{
		vector: vector,
	}
}

func (v *Vector) Add(item interface{}) {
	if v.length >= v.capacity {
		v.resize()
	}

	v.items[v.length-1] = item
}

func (v *Vector) Set(item interface{}, index int) bool {
	if index > v.length-1 {
		return false
	}

	v.items[index] = item

	return true
}

func (v *Vector) Get(index int) interface{} {
	val, _ := v.GetExisted(index)
	return val
}

func (v *Vector) GetExisted(index int) (interface{}, bool) {
	if index > v.length-1 {
		return nil, false
	}

	return v.items[index], true
}

func (v *Vector) Len() int {
	return v.length
}

func (v *Vector) Cap() int {
	return v.capacity
}

func (v *Vector) Remove(index int) (interface{}, bool) {
	if index > v.length-1 {
		return nil, false
	}

	val := v.items[index]

	for i := index; i < v.length-1; i++ {
		tmp := v.items[i]
		v.items[i] = v.items[i+1]
		v.items[i+1] = tmp
	}

	v.length--

	return val, true
}

func (v *Vector) resize() {
	items := make([]interface{}, v.length+v.vector)
	copy(items, v.items)
	v.items = items
	v.capacity++
	v.length++
}
