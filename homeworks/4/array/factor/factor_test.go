package factor

import (
	"reflect"
	"testing"
)

func TestFactor_Append(t *testing.T) {
	factor := NewFactor()

	if factor.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", factor.Len(), 0)
	}

	factor.Append(1)

	if factor.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", factor.Len(), 1)
	}

	factor.Append(1)

	if factor.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", factor.Len(), 2)
	}
}

func TestFactor_SetGet(t *testing.T) {
	factor := NewFactor()
	factor.Append(1)
	factor.Append(2)
	factor.Append(3)

	var ok bool

	ok = factor.Set(10, 0)
	if !ok {
		t.Fatalf("unexpected not ok Set(10, 0)")
	}

	ok = factor.Set(20, 1)
	if !ok {
		t.Fatalf("unexpected not ok Set(20, 1)")
	}

	factor.Set(30, 2)

	if !ok {
		t.Fatalf("unexpected not ok Set(30, 2)")
	}

	if factor.Get(0) != 10 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", factor.Get(0), 0, 10)
	}

	if factor.Get(1) != 20 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", factor.Get(1), 1, 20)
	}

	if factor.Get(2) != 30 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", factor.Get(2), 2, 30)
	}

	ok = factor.Set(100, 10)
	if ok {
		t.Fatalf("unexpected ok Set(100, 10)")
	}

	val := factor.Get(10)
	if val != nil {
		t.Fatalf("unexpected value %v got by index instead of nil", val)
	}
}

func TestFactor_GetExisted(t *testing.T) {
	factor := NewFactor()
	factor.Append(1)

	val, ok := factor.GetExisted(0)

	if !ok {
		t.Fatalf("unexpected not ok GetExisted(0)")
	}

	if val != 1 {
		t.Fatalf("unexpected value %d got by index %d instread of %d", val, 0, 1)
	}

	_, ok = factor.GetExisted(10)
	if ok {
		t.Fatalf("unexpected ok GetExisted(10)")
	}
}

func TestFactor_Remove(t *testing.T) {
	factor := NewFactor()

	_, ok := factor.Remove(10)
	if ok {
		t.Fatalf("unexpected ok removing from empty array")
	}

	factor.Append(1) // x
	factor.Append(2)
	factor.Append(3) // x
	factor.Append(4)
	factor.Append(5)
	factor.Append(6)
	factor.Append(7) // x

	val, ok := factor.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	val, ok = factor.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	val, ok = factor.Remove(factor.Len() - 1)

	if !ok {
		t.Fatalf("unexpected not ok removing last item")
	}

	if val != 7 {
		t.Fatalf("unexpected value %d of removed last item instread of %d", val, 7)
	}

	if factor.Len() != 4 {
		t.Fatalf("unexpected length %d of array after three removals instead of %d", factor.Len(), 4)
	}

	if factor.Cap() != 7 {
		t.Fatalf("unexpected capacity %d of array after three removals instead of %d", factor.Cap(), 7)
	}

	slice := convertToSlice(factor)
	expected := []interface{}{2, 4, 5, 6}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestFactor_AddAtTheEnd(t *testing.T) {
	factor := NewFactor()

	factor.Append(1)
	factor.Append(2)
	factor.Append(3)
	factor.Append(4)
	factor.Append(5)
	factor.Append(6)
	factor.Append(7)

	ok := factor.Add(10, 10)
	if ok {
		t.Fatalf("unexpected ok adding item out of range")
	}

	// same as append
	ok = factor.Add(100, factor.Len())

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(factor)
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 100}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestFactor_AddAtTheMiddle(t *testing.T) {
	factor := NewFactor()

	factor.Append(1)
	factor.Append(2)
	factor.Append(3)
	factor.Append(4)
	factor.Append(5)
	factor.Append(6)
	factor.Append(7)

	// same as append
	ok := factor.Add(100, 2)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(factor)
	expected := []interface{}{1, 2, 100, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestFactor_AddAtTheBeginning(t *testing.T) {
	factor := NewFactor()

	factor.Append(1)
	factor.Append(2)
	factor.Append(3)
	factor.Append(4)
	factor.Append(5)
	factor.Append(6)
	factor.Append(7)

	// same as append
	ok := factor.Add(100, 0)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(factor)
	expected := []interface{}{100, 1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

func TestFactor_AddAfterRemoval(t *testing.T) {
	factor := NewFactor()

	factor.Append(1)
	factor.Append(2)
	factor.Append(3)
	factor.Append(4)
	factor.Append(5) // x
	factor.Append(6)
	factor.Append(7)
	factor.Append(8) // x
	factor.Append(9)
	factor.Append(10)

	factor.Remove(4) // x=5
	factor.Remove(6) // x=8

	// same as append
	ok := factor.Add(100, 5)

	if !ok {
		t.Fatalf("unexpected not ok adding item in the end")
	}

	slice := convertToSlice(factor)
	expected := []interface{}{1, 2, 3, 4, 6, 100, 7, 9, 10}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

// convertToSlice converts our implementation of array to GO slice - helper to tests.
func convertToSlice(f *Factor) []interface{} {
	slice := make([]interface{}, f.Len())
	copy(slice, f.items)

	return slice
}
