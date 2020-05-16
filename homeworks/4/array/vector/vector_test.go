package vector

import "testing"

func TestVector_AddLen(t *testing.T) {
	vector := NewVector()

	if vector.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", vector.Len(), 0)
	}

	vector.Add(1)

	if vector.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", vector.Len(), 1)
	}

	vector.Add(1)

	if vector.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", vector.Len(), 2)
	}
}

func TestVector_SetGet(t *testing.T) {
	vector := NewVector()
	vector.Add(1)
	vector.Add(2)
	vector.Add(3)

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
	vector.Add(1)

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

	vector.Add(1) // x
	vector.Add(2)
	vector.Add(3) // x
	vector.Add(4)
	vector.Add(5)
	vector.Add(6)
	vector.Add(7) // x

	val, ok := vector.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	if vector.Len() != 6 {
		t.Fatalf("unexpected length %d of array after one removal instead of %d", vector.Len(), 6)
	}

	val, ok = vector.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	if vector.Len() != 5 {
		t.Fatalf("unexpected length %d of array after two removals instead of %d", vector.Len(), 5)
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
}
