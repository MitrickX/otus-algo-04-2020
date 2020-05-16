package list

import (
	"reflect"
	"testing"
)

func TestList_Append(t *testing.T) {
	list := NewList()

	if list.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", list.Len(), 0)
	}

	list.Append(1)

	if list.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", list.Len(), 1)
	}

	list.Append(1)

	if list.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", list.Len(), 2)
	}
}

func TestList_SetGet(t *testing.T) {
	list := NewList()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	var ok bool

	ok = list.Set(10, 0)
	if !ok {
		t.Fatalf("unexpected not ok Set(10, 0)")
	}

	ok = list.Set(20, 1)
	if !ok {
		t.Fatalf("unexpected not ok Set(20, 1)")
	}

	list.Set(30, 2)

	if !ok {
		t.Fatalf("unexpected not ok Set(30, 2)")
	}

	if list.Get(0) != 10 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", list.Get(0), 0, 10)
	}

	if list.Get(1) != 20 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", list.Get(1), 1, 20)
	}

	if list.Get(2) != 30 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", list.Get(2), 2, 30)
	}

	ok = list.Set(100, 10)
	if ok {
		t.Fatalf("unexpected ok Set(100, 10)")
	}

	val := list.Get(10)
	if val != nil {
		t.Fatalf("unexpected value %v got by index instead of nil", val)
	}
}

func TestList_GetExisted(t *testing.T) {
	list := NewList()
	list.Append(1)

	val, ok := list.GetExisted(0)

	if !ok {
		t.Fatalf("unexpected not ok GetExisted(0)")
	}

	if val != 1 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", val, 0, 1)
	}

	_, ok = list.GetExisted(10)
	if ok {
		t.Fatalf("unexpected ok GetExisted(10)")
	}
}

func TestList_Remove(t *testing.T) {
	list := NewList()

	_, ok := list.Remove(10)
	if ok {
		t.Fatalf("unexpected ok removing from empty array")
	}

	list.Append(1) // x
	list.Append(2)
	list.Append(3) // x
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7) // x

	val, ok := list.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	val, ok = list.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	val, ok = list.Remove(list.Len() - 1)

	if !ok {
		t.Fatalf("unexpected not ok removing last item")
	}

	if val != 7 {
		t.Fatalf("unexpected value %d of removed last item instread of %d", val, 7)
	}

	if list.Len() != 4 {
		t.Fatalf("unexpected length %d of array after three removals instead of %d", list.Len(), 4)
	}

	slice := convertToSlice(list)
	expected := []interface{}{2, 4, 5, 6}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestList_AddAtTheEnd(t *testing.T) {
	list := NewList()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)

	ok := list.Add(10, 10)
	if ok {
		t.Fatalf("unexpected ok adding item out of range")
	}

	// same as append
	ok = list.Add(100, list.Len())

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(list)
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 100}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestList_AddAtTheMiddle(t *testing.T) {
	list := NewList()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)

	// same as append
	ok := list.Add(100, 2)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(list)
	expected := []interface{}{1, 2, 100, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestList_AddAtTheBeginning(t *testing.T) {
	list := NewList()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)

	// same as append
	ok := list.Add(100, 0)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(list)
	expected := []interface{}{100, 1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestList_AddAfterRemoval(t *testing.T) {
	list := NewList()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5) // x
	list.Append(6)
	list.Append(7)
	list.Append(8) // x
	list.Append(9)
	list.Append(10)

	list.Remove(4) // x=5
	list.Remove(6) // x=8

	// same as append
	ok := list.Add(100, 5)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(list)
	expected := []interface{}{1, 2, 3, 4, 6, 100, 7, 9, 10}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

// convertToSlice converts our implementation of array to GO slice - helper to tests.
func convertToSlice(l *List) []interface{} {
	slice := make([]interface{}, l.Len())
	copy(slice, l.items)

	return slice
}
