package matrix

import (
	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/container"
	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/container/factor"
)

type Array struct {
	vector int
	length int
	array  container.Array
}

func NewArray() *Array {
	return NewArrayCustom(100)
}

func NewArrayCustom(vector int) *Array {
	return &Array{
		vector: vector,
		array:  factor.NewArray(),
	}
}

func (a *Array) Add(item interface{}, index int) bool {
	if index > a.length {
		return false
	}

	// append at the end and then bubble value from end up to index
	a.Append(item)

	if index == a.length-1 {
		return true
	}

	// index of row in table
	var rowIndex int

	// slice of current row
	var currentSlice []interface{}

	for i := a.length - 1; i > index; i-- {
		// all this crazy code for making runtime assertion one time per row (not per loop step)
		if currentSlice == nil || i/a.vector != rowIndex {
			rowIndex = i / a.vector
			row := a.array.Get(rowIndex)
			currentSlice = row.([]interface{}) // runtime interface{} assertion
		}

		colIndex := i % a.vector

		// crazy swap with taking into account boundary case (when col == 0)
		// in this case we must swap with last item in adjacentSlice

		if colIndex == 0 {
			row := a.array.Get(rowIndex - 1)
			adjacentSlice := row.([]interface{}) // runtime interface{} assertion
			tmp := currentSlice[colIndex]
			currentSlice[colIndex] = adjacentSlice[a.vector-1]
			adjacentSlice[a.vector-1] = tmp
		} else {
			tmp := currentSlice[colIndex]
			currentSlice[colIndex] = currentSlice[colIndex-1]
			currentSlice[colIndex-1] = tmp
		}
	}

	return true
}

func (a *Array) Append(item interface{}) {
	if a.length == a.array.Len()*a.vector {
		a.resize()
	}

	row := a.array.Get(a.length / a.vector)

	// runtime interface{} assertion
	slice := row.([]interface{})

	slice[a.length%a.vector] = item
	a.length++
}

func (a *Array) Set(item interface{}, index int) bool {
	if index > a.length-1 {
		return false
	}

	row := a.array.Get(index / a.vector)

	// runtime interface{} assertion
	slice := row.([]interface{})

	slice[index%a.vector] = item

	return true
}

func (a *Array) Get(index int) interface{} {
	val, ok := a.GetExisted(index)
	if !ok {
		return nil
	}

	return val
}

func (a *Array) GetExisted(index int) (interface{}, bool) {
	if index > a.length-1 {
		return nil, false
	}

	row, ok := a.array.GetExisted(index / a.vector)
	if !ok {
		return nil, false
	}

	// runtime interface{} assertion
	slice, ok := row.([]interface{})

	if !ok {
		return nil, false
	}

	return slice[index%a.vector], true
}

func (a *Array) Len() int {
	return a.length
}

func (a *Array) Cap() int {
	return a.array.Cap() * a.vector
}

func (a *Array) Remove(index int) (interface{}, bool) {
	if index > a.length-1 {
		return nil, false
	}

	val := a.Get(index)

	// index of row in table
	var rowIndex int

	// slice of current row
	var currentSlice []interface{}

	for i := index; i < a.length-1; i++ {
		// all this crazy code for making runtime assertion one time per row (not per loop step)
		if currentSlice == nil || i/a.vector != rowIndex {
			rowIndex = i / a.vector
			row := a.array.Get(rowIndex)
			currentSlice = row.([]interface{}) // runtime interface{} assertion
		}

		colIndex := i % a.vector

		// crazy swap with taking into account boundary case (when col == vector-1)
		// in this case we must swap with first item in adjacentSlice

		if colIndex == a.vector-1 {
			row := a.array.Get(rowIndex + 1)
			adjacentSlice := row.([]interface{}) // runtime interface{} assertion
			tmp := currentSlice[colIndex]
			currentSlice[colIndex] = adjacentSlice[0]
			adjacentSlice[0] = tmp
		} else {
			tmp := currentSlice[colIndex]
			currentSlice[colIndex] = currentSlice[colIndex+1]
			currentSlice[colIndex+1] = tmp
		}
	}

	a.length--

	return val, true
}

func (a *Array) resize() {
	a.array.Append(make([]interface{}, a.vector))
}
