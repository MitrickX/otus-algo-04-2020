package vector

import (
	"reflect"
	"testing"
)

func TestVector_Append(t *testing.T) {
	vector := NewVector()

	if vector.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", vector.Len(), 0)
	}

	vector.Append(1)

	if vector.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", vector.Len(), 1)
	}

	vector.Append(1)

	if vector.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", vector.Len(), 2)
	}
}

func TestVector_SetGet(t *testing.T) {
	vector := NewVector()
	vector.Append(1)
	vector.Append(2)
	vector.Append(3)

	var ok bool

	ok = vector.Set(10, 0)
	if !ok {
		t.Fatalf("unexpected not ok Set(10, 0)")
	}

	ok = vector.Set(20, 1)
	if !ok {
		t.Fatalf("unexpected not ok Set(20, 1)")
	}

	vector.Set(30, 2)

	if !ok {
		t.Fatalf("unexpected not ok Set(30, 2)")
	}

	if vector.Get(0) != 10 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", vector.Get(0), 0, 10)
	}

	if vector.Get(1) != 20 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", vector.Get(1), 1, 20)
	}

	if vector.Get(2) != 30 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", vector.Get(2), 2, 30)
	}

	ok = vector.Set(100, 10)
	if ok {
		t.Fatalf("unexpected ok Set(100, 10)")
	}

	val := vector.Get(10)
	if val != nil {
		t.Fatalf("unexpected value %v got by index instead of nil", val)
	}
}

func TestVector_GetExisted(t *testing.T) {
	vector := NewVector()
	vector.Append(1)

	val, ok := vector.GetExisted(0)

	if !ok {
		t.Fatalf("unexpected not ok GetExisted(0)")
	}

	if val != 1 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", val, 0, 1)
	}

	_, ok = vector.GetExisted(10)
	if ok {
		t.Fatalf("unexpected ok GetExisted(10)")
	}
}

func TestVector_Remove(t *testing.T) {
	vector := NewVector()

	_, ok := vector.Remove(10)
	if ok {
		t.Fatalf("unexpected ok removing from empty array")
	}

	vector.Append(1) // x
	vector.Append(2)
	vector.Append(3) // x
	vector.Append(4)
	vector.Append(5)
	vector.Append(6)
	vector.Append(7) // x

	val, ok := vector.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	val, ok = vector.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	val, ok = vector.Remove(vector.Len() - 1)

	if !ok {
		t.Fatalf("unexpected not ok removing last item")
	}

	if val != 7 {
		t.Fatalf("unexpected value %d of removed last item instread of %d", val, 7)
	}

	if vector.Len() != 4 {
		t.Fatalf("unexpected length %d of array after three removals instead of %d", vector.Len(), 4)
	}

	if vector.Cap() != 7 {
		t.Fatalf("unexpected capacity %d of array after three removals instead of %d", vector.Cap(), 7)
	}

	slice := convertToSlice(vector)
	expected := []interface{}{2, 4, 5, 6}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestVector_AddAtTheEnd(t *testing.T) {
	vector := NewVector()

	vector.Append(1)
	vector.Append(2)
	vector.Append(3)
	vector.Append(4)
	vector.Append(5)
	vector.Append(6)
	vector.Append(7)

	ok := vector.Add(10, 10)
	if ok {
		t.Fatalf("unexpected ok adding item out of range")
	}

	// same as append
	ok = vector.Add(100, vector.Len())

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(vector)
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 100}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestVector_AddAtTheMiddle(t *testing.T) {
	vector := NewVector()

	vector.Append(1)
	vector.Append(2)
	vector.Append(3)
	vector.Append(4)
	vector.Append(5)
	vector.Append(6)
	vector.Append(7)

	// same as append
	ok := vector.Add(100, 2)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(vector)
	expected := []interface{}{1, 2, 100, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestVector_AddAtTheBeginning(t *testing.T) {
	vector := NewVector()

	vector.Append(1)
	vector.Append(2)
	vector.Append(3)
	vector.Append(4)
	vector.Append(5)
	vector.Append(6)
	vector.Append(7)

	// same as append
	ok := vector.Add(100, 0)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(vector)
	expected := []interface{}{100, 1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestVector_AddAfterRemoval(t *testing.T) {
	vector := NewVector()

	vector.Append(1)
	vector.Append(2)
	vector.Append(3)
	vector.Append(4)
	vector.Append(5) // x
	vector.Append(6)
	vector.Append(7)
	vector.Append(8) // x
	vector.Append(9)
	vector.Append(10)

	vector.Remove(4) // x=5
	vector.Remove(6) // x=8

	// same as append
	ok := vector.Add(100, 5)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(vector)
	expected := []interface{}{1, 2, 3, 4, 6, 100, 7, 9, 10}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

// convertToSlice converts our implementation of array to GO slice - helper to tests.
func convertToSlice(v *Vector) []interface{} {
	slice := make([]interface{}, v.Len())
	copy(slice, v.items)

	return slice
}
