package matrix

import (
	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/array"
	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/array/factor"
)

type Matrix struct {
	vector int
	length int
	array  array.Array
}

func NewMatrix() *Matrix {
	return NewMatrixCustom(100)
}

func NewMatrixCustom(vector int) *Matrix {
	return &Matrix{
		vector: vector,
		array:  factor.NewFactor(),
	}
}

func (m *Matrix) Add(item interface{}, index int) bool {
	if index > m.length {
		return false
	}

	// append at the end and then bubble value from end up to index
	m.Append(item)

	if index == m.length-1 {
		return true
	}

	// index of row in table
	var rowIndex int

	// slice of current row
	var currentSlice []interface{}

	for i := m.length - 1; i > index; i-- {
		// all this crazy code for making runtime assertion one time per row (not per loop step)
		if currentSlice == nil || i/m.vector != rowIndex {
			rowIndex = i / m.vector
			row := m.array.Get(rowIndex)
			currentSlice = row.([]interface{}) // runtime interface{} assertion
		}

		colIndex := i % m.vector

		// crazy swap with taking into account boundary case (when col == 0)
		// in this case we must swap with last item in adjacentSlice

		if colIndex == 0 {
			row := m.array.Get(rowIndex - 1)
			adjacentSlice := row.([]interface{}) // runtime interface{} assertion
			tmp := currentSlice[colIndex]
			currentSlice[colIndex] = adjacentSlice[m.vector-1]
			adjacentSlice[m.vector-1] = tmp
		} else {
			tmp := currentSlice[colIndex]
			currentSlice[colIndex] = currentSlice[colIndex-1]
			currentSlice[colIndex-1] = tmp
		}
	}

	return true
}

func (m *Matrix) Append(item interface{}) {
	if m.length == m.array.Len()*m.vector {
		m.resize()
	}

	row := m.array.Get(m.length / m.vector)

	// runtime interface{} assertion
	slice := row.([]interface{})

	slice[m.length%m.vector] = item
	m.length++
}

func (m *Matrix) Set(item interface{}, index int) bool {
	if index > m.length-1 {
		return false
	}

	row := m.array.Get(index / m.vector)

	// runtime interface{} assertion
	slice := row.([]interface{})

	slice[index%m.vector] = item

	return true
}

func (m *Matrix) Get(index int) interface{} {
	val, ok := m.GetExisted(index)
	if !ok {
		return nil
	}

	return val
}

func (m *Matrix) GetExisted(index int) (interface{}, bool) {
	if index > m.length-1 {
		return nil, false
	}

	row, ok := m.array.GetExisted(index / m.vector)
	if !ok {
		return nil, false
	}

	// runtime interface{} assertion
	slice, ok := row.([]interface{})

	if !ok {
		return nil, false
	}

	return slice[index%m.vector], true
}

func (m *Matrix) Len() int {
	return m.length
}

func (m *Matrix) Cap() int {
	return m.array.Cap() * m.vector
}

func (m *Matrix) Remove(index int) (interface{}, bool) {
	if index > m.length-1 {
		return nil, false
	}

	val := m.Get(index)

	// index of row in table
	var rowIndex int

	// slice of current row
	var currentSlice []interface{}

	for i := index; i < m.length-1; i++ {
		// all this crazy code for making runtime assertion one time per row (not per loop step)
		if currentSlice == nil || i/m.vector != rowIndex {
			rowIndex = i / m.vector
			row := m.array.Get(rowIndex)
			currentSlice = row.([]interface{}) // runtime interface{} assertion
		}

		colIndex := i % m.vector

		// crazy swap with taking into account boundary case (when col == vector-1)
		// in this case we must swap with first item in adjacentSlice

		if colIndex == m.vector-1 {
			row := m.array.Get(rowIndex + 1)
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

	m.length--

	return val, true
}

func (m *Matrix) resize() {
	m.array.Append(make([]interface{}, m.vector))
}
