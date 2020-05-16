package matrix

import (
	"reflect"
	"testing"
)

func TestMatrix_Append(t *testing.T) {
	matrix := NewMatrixCustom(5)

	if matrix.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", matrix.Len(), 0)
	}

	matrix.Append(1)

	if matrix.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", matrix.Len(), 1)
	}

	matrix.Append(1)

	if matrix.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", matrix.Len(), 2)
	}
}

func TestMatrix_SetGet(t *testing.T) {
	matrix := NewMatrixCustom(5)
	matrix.Append(1)
	matrix.Append(2)
	matrix.Append(3)

	var ok bool

	ok = matrix.Set(10, 0)
	if !ok {
		t.Fatalf("unexpected not ok Set(10, 0)")
	}

	ok = matrix.Set(20, 1)
	if !ok {
		t.Fatalf("unexpected not ok Set(20, 1)")
	}

	matrix.Set(30, 2)

	if !ok {
		t.Fatalf("unexpected not ok Set(30, 2)")
	}

	if matrix.Get(0) != 10 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", matrix.Get(0), 0, 10)
	}

	if matrix.Get(1) != 20 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", matrix.Get(1), 1, 20)
	}

	if matrix.Get(2) != 30 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", matrix.Get(2), 2, 30)
	}

	ok = matrix.Set(100, 10)
	if ok {
		t.Fatalf("unexpected ok Set(100, 10)")
	}

	val := matrix.Get(10)
	if val != nil {
		t.Fatalf("unexpected value %v got by index instead of nil", val)
	}
}

func TestMatrix_GetExisted(t *testing.T) {
	matrix := NewMatrixCustom(5)
	matrix.Append(1)

	val, ok := matrix.GetExisted(0)

	if !ok {
		t.Fatalf("unexpected not ok GetExisted(0)")
	}

	if val != 1 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", val, 0, 1)
	}

	_, ok = matrix.GetExisted(1)
	if ok {
		t.Fatalf("unexpected ok GetExisted(1)")
	}

	_, ok = matrix.GetExisted(10)
	if ok {
		t.Fatalf("unexpected ok GetExisted(10)")
	}
}

func TestMatrix_Remove(t *testing.T) {
	matrix := NewMatrixCustom(5)

	_, ok := matrix.Remove(10)
	if ok {
		t.Fatalf("unexpected ok removing from empty array")
	}

	matrix.Append(1) // x
	matrix.Append(2)
	matrix.Append(3) // x
	matrix.Append(4)
	matrix.Append(5)
	matrix.Append(6)
	matrix.Append(7) // x

	val, ok := matrix.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	val, ok = matrix.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	val, ok = matrix.Remove(matrix.Len() - 1)

	if !ok {
		t.Fatalf("unexpected not ok removing last item")
	}

	if val != 7 {
		t.Fatalf("unexpected value %d of removed last item instread of %d", val, 7)
	}

	if matrix.Len() != 4 {
		t.Fatalf("unexpected length %d of array after three removals instead of %d", matrix.Len(), 4)
	}

	slice := convertToSlice(matrix)
	expected := []interface{}{2, 4, 5, 6}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestMatrix_AddAtTheEnd(t *testing.T) {
	matrix := NewMatrixCustom(5)

	matrix.Append(1)
	matrix.Append(2)
	matrix.Append(3)
	matrix.Append(4)
	matrix.Append(5)
	matrix.Append(6)
	matrix.Append(7)

	ok := matrix.Add(10, 10)
	if ok {
		t.Fatalf("unexpected ok adding item out of range")
	}

	// same as append
	ok = matrix.Add(100, matrix.Len())

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(matrix)
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 100}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestMatrix_AddAtTheMiddle(t *testing.T) {
	matrix := NewMatrixCustom(5)

	matrix.Append(1)
	matrix.Append(2)
	matrix.Append(3)
	matrix.Append(4)
	matrix.Append(5)
	matrix.Append(6)
	matrix.Append(7)

	// same as append
	ok := matrix.Add(100, 2)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(matrix)
	expected := []interface{}{1, 2, 100, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestMatrix_AddAtTheBeginning(t *testing.T) {
	matrix := NewMatrixCustom(5)

	matrix.Append(1)
	matrix.Append(2)
	matrix.Append(3)
	matrix.Append(4)
	matrix.Append(5)
	matrix.Append(6)
	matrix.Append(7)

	// same as append
	ok := matrix.Add(100, 0)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(matrix)
	expected := []interface{}{100, 1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestMatrix_AddAfterRemoval(t *testing.T) {
	matrix := NewMatrixCustom(5)

	matrix.Append(1)
	matrix.Append(2)
	matrix.Append(3)
	matrix.Append(4)
	matrix.Append(5) // x
	matrix.Append(6)
	matrix.Append(7)
	matrix.Append(8) // x
	matrix.Append(9)
	matrix.Append(10)

	matrix.Remove(4) // x=5
	matrix.Remove(6) // x=8

	// same as append
	ok := matrix.Add(100, 5)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(matrix)
	expected := []interface{}{1, 2, 3, 4, 6, 100, 7, 9, 10}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

// convertToSlice converts our implementation of array to GO slice - helper to tests.
func convertToSlice(m *Matrix) []interface{} {
	slice := make([]interface{}, m.Len())
	arLen := m.Len()

	for i := 0; i < arLen; i++ {
		slice[i] = m.Get(i)
	}

	return slice
}
