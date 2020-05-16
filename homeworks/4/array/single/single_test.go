package single

import (
	"reflect"
	"testing"
)

func TestSingle_Append(t *testing.T) {
	single := NewSingle()

	if single.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", single.Len(), 0)
	}

	single.Append(1)

	if single.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", single.Len(), 1)
	}

	single.Append(1)

	if single.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", single.Len(), 2)
	}
}

func TestSingle_SetGet(t *testing.T) {
	single := NewSingle()
	single.Append(1)
	single.Append(2)
	single.Append(3)

	var ok bool

	ok = single.Set(10, 0)
	if !ok {
		t.Fatalf("unexpected not ok Set(10, 0)")
	}

	ok = single.Set(20, 1)
	if !ok {
		t.Fatalf("unexpected not ok Set(20, 1)")
	}

	single.Set(30, 2)

	if !ok {
		t.Fatalf("unexpected not ok Set(30, 2)")
	}

	if single.Get(0) != 10 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", single.Get(0), 0, 10)
	}

	if single.Get(1) != 20 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", single.Get(1), 1, 20)
	}

	if single.Get(2) != 30 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", single.Get(2), 2, 30)
	}

	ok = single.Set(100, 10)
	if ok {
		t.Fatalf("unexpected ok Set(100, 10)")
	}

	val := single.Get(10)
	if val != nil {
		t.Fatalf("unexpected value %v got by index instead of nil", val)
	}
}

func TestSingle_GetExisted(t *testing.T) {
	single := NewSingle()
	single.Append(1)

	val, ok := single.GetExisted(0)

	if !ok {
		t.Fatalf("unexpected not ok GetExisted(0)")
	}

	if val != 1 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", val, 0, 1)
	}

	_, ok = single.GetExisted(10)
	if ok {
		t.Fatalf("unexpected ok GetExisted(10)")
	}
}

func TestSingle_Remove(t *testing.T) {
	single := NewSingle()

	_, ok := single.Remove(10)
	if ok {
		t.Fatalf("unexpected ok removing from empty array")
	}

	single.Append(1) // x
	single.Append(2)
	single.Append(3) // x
	single.Append(4)
	single.Append(5)
	single.Append(6)
	single.Append(7) // x

	val, ok := single.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	val, ok = single.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	val, ok = single.Remove(single.Len() - 1)

	if !ok {
		t.Fatalf("unexpected not ok removing last item")
	}

	if val != 7 {
		t.Fatalf("unexpected value %d of removed last item instread of %d", val, 7)
	}

	if single.Len() != 4 {
		t.Fatalf("unexpected length %d of array after three removals instead of %d", single.Len(), 4)
	}

	if single.Cap() != 5 {
		t.Fatalf("unexpected capacity %d of array after three removals instead of %d", single.Cap(), 5)
	}

	slice := convertToSlice(single)
	expected := []interface{}{2, 4, 5, 6}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestSingle_AddAtTheEnd(t *testing.T) {
	single := NewSingle()

	single.Append(1)
	single.Append(2)
	single.Append(3)
	single.Append(4)
	single.Append(5)
	single.Append(6)
	single.Append(7)

	ok := single.Add(10, 10)
	if ok {
		t.Fatalf("unexpected ok adding item out of range")
	}

	// same as append
	ok = single.Add(100, single.Len())

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(single)
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 100}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestSingle_AddAtTheMiddle(t *testing.T) {
	single := NewSingle()

	single.Append(1)
	single.Append(2)
	single.Append(3)
	single.Append(4)
	single.Append(5)
	single.Append(6)
	single.Append(7)

	// same as append
	ok := single.Add(100, 2)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(single)
	expected := []interface{}{1, 2, 100, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestSingle_AddAtTheBeginning(t *testing.T) {
	single := NewSingle()

	single.Append(1)
	single.Append(2)
	single.Append(3)
	single.Append(4)
	single.Append(5)
	single.Append(6)
	single.Append(7)

	// same as append
	ok := single.Add(100, 0)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(single)
	expected := []interface{}{100, 1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestSingle_AddAfterRemoval(t *testing.T) {
	single := NewSingle()

	single.Append(1)
	single.Append(2)
	single.Append(3)
	single.Append(4)
	single.Append(5) // x
	single.Append(6)
	single.Append(7)
	single.Append(8) // x
	single.Append(9)
	single.Append(10)

	single.Remove(4) // x=5
	single.Remove(6) // x=8

	// same as append
	ok := single.Add(100, 5)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(single)
	expected := []interface{}{1, 2, 3, 4, 6, 100, 7, 9, 10}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

// convertToSlice converts our implementation of array to GO slice - helper to tests.
func convertToSlice(s *Single) []interface{} {
	slice := make([]interface{}, s.Len())
	copy(slice, s.items)

	return slice
}
