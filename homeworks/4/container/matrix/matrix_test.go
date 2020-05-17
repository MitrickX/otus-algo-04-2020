package matrix

import (
	"reflect"
	"testing"
)

func TestArray_Append(t *testing.T) {
	array := NewArrayCustom(5)

	if array.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", array.Len(), 0)
	}

	array.Append(1)

	if array.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", array.Len(), 1)
	}

	array.Append(1)

	if array.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", array.Len(), 2)
	}
}

func TestArray_SetGet(t *testing.T) {
	array := NewArrayCustom(5)
	array.Append(1)
	array.Append(2)
	array.Append(3)

	var ok bool

	ok = array.Set(10, 0)
	if !ok {
		t.Fatalf("unexpected not ok Set(10, 0)")
	}

	ok = array.Set(20, 1)
	if !ok {
		t.Fatalf("unexpected not ok Set(20, 1)")
	}

	array.Set(30, 2)

	if !ok {
		t.Fatalf("unexpected not ok Set(30, 2)")
	}

	if array.Get(0) != 10 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", array.Get(0), 0, 10)
	}

	if array.Get(1) != 20 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", array.Get(1), 1, 20)
	}

	if array.Get(2) != 30 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", array.Get(2), 2, 30)
	}

	ok = array.Set(100, 10)
	if ok {
		t.Fatalf("unexpected ok Set(100, 10)")
	}

	val := array.Get(10)
	if val != nil {
		t.Fatalf("unexpected value %v got by index instead of nil", val)
	}
}

func TestArray_GetExisted(t *testing.T) {
	array := NewArrayCustom(5)
	array.Append(1)

	val, ok := array.GetExisted(0)

	if !ok {
		t.Fatalf("unexpected not ok GetExisted(0)")
	}

	if val != 1 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", val, 0, 1)
	}

	_, ok = array.GetExisted(1)
	if ok {
		t.Fatalf("unexpected ok GetExisted(1)")
	}

	_, ok = array.GetExisted(10)
	if ok {
		t.Fatalf("unexpected ok GetExisted(10)")
	}
}

func TestArray_Remove(t *testing.T) {
	array := NewArrayCustom(5)

	_, ok := array.Remove(10)
	if ok {
		t.Fatalf("unexpected ok removing from empty array")
	}

	array.Append(1) // x
	array.Append(2)
	array.Append(3) // x
	array.Append(4)
	array.Append(5)
	array.Append(6)
	array.Append(7) // x

	val, ok := array.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	val, ok = array.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	val, ok = array.Remove(array.Len() - 1)

	if !ok {
		t.Fatalf("unexpected not ok removing last item")
	}

	if val != 7 {
		t.Fatalf("unexpected value %d of removed last item instread of %d", val, 7)
	}

	if array.Len() != 4 {
		t.Fatalf("unexpected length %d of array after three removals instead of %d", array.Len(), 4)
	}

	slice := convertToSlice(array)
	expected := []interface{}{2, 4, 5, 6}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestArray_AddAtTheEnd(t *testing.T) {
	array := NewArrayCustom(5)

	array.Append(1)
	array.Append(2)
	array.Append(3)
	array.Append(4)
	array.Append(5)
	array.Append(6)
	array.Append(7)

	ok := array.Add(10, 10)
	if ok {
		t.Fatalf("unexpected ok adding item out of range")
	}

	// same as append
	ok = array.Add(100, array.Len())

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(array)
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 100}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestArray_AddAtTheMiddle(t *testing.T) {
	array := NewArrayCustom(5)

	array.Append(1)
	array.Append(2)
	array.Append(3)
	array.Append(4)
	array.Append(5)
	array.Append(6)
	array.Append(7)

	// same as append
	ok := array.Add(100, 2)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(array)
	expected := []interface{}{1, 2, 100, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestArray_AddAtTheBeginning(t *testing.T) {
	array := NewArrayCustom(5)

	array.Append(1)
	array.Append(2)
	array.Append(3)
	array.Append(4)
	array.Append(5)
	array.Append(6)
	array.Append(7)

	// same as append
	ok := array.Add(100, 0)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(array)
	expected := []interface{}{100, 1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestArray_AddAfterRemoval(t *testing.T) {
	array := NewArrayCustom(5)

	array.Append(1)
	array.Append(2)
	array.Append(3)
	array.Append(4)
	array.Append(5) // x
	array.Append(6)
	array.Append(7)
	array.Append(8) // x
	array.Append(9)
	array.Append(10)

	array.Remove(4) // x=5
	array.Remove(6) // x=8

	// same as append
	ok := array.Add(100, 5)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(array)
	expected := []interface{}{1, 2, 3, 4, 6, 100, 7, 9, 10}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

// convertToSlice converts our implementation of array to GO slice - helper to tests.
func convertToSlice(a *Array) []interface{} {
	slice := make([]interface{}, a.Len())
	arLen := a.Len()

	for i := 0; i < arLen; i++ {
		slice[i] = a.Get(i)
	}

	return slice
}
