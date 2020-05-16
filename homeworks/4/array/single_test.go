package array

import (
	"testing"
)

func TestSingle_AddLen(t *testing.T) {
	single := NewSingle()

	if single.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", single.Len(), 0)
	}

	single.Add(1)

	if single.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", single.Len(), 1)
	}

	single.Add(1)

	if single.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", single.Len(), 2)
	}
}

func TestSingle_SetGet(t *testing.T) {
	single := NewSingle()
	single.Add(1)
	single.Add(2)
	single.Add(3)

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
	single.Add(1)

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

	single.Add(1) // x
	single.Add(2)
	single.Add(3) // x
	single.Add(4)
	single.Add(5)
	single.Add(6)
	single.Add(7) // x

	val, ok := single.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	if single.Len() != 6 {
		t.Fatalf("unexpected length %d of array after one removal instead of %d", single.Len(), 6)
	}

	val, ok = single.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	if single.Len() != 5 {
		t.Fatalf("unexpected length %d of array after two removals instead of %d", single.Len(), 5)
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

	if single.Cap() != 7 {
		t.Fatalf("unexpected capacity %d of array after three removals instead of %d", single.Cap(), 7)
	}
}
