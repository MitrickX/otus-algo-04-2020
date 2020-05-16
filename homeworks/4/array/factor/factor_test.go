package factor

import "testing"

func TestFactor_AddLen(t *testing.T) {
	factor := NewFactor()

	if factor.Len() != 0 {
		t.Fatalf("unexpected len %d instead of %d", factor.Len(), 0)
	}

	factor.Add(1)

	if factor.Len() != 1 {
		t.Fatalf("unexpected len %d instead of %d", factor.Len(), 1)
	}

	factor.Add(1)

	if factor.Len() != 2 {
		t.Fatalf("unexpected len %d instead of %d", factor.Len(), 2)
	}
}

func TestFactor_SetGet(t *testing.T) {
	factor := NewFactor()
	factor.Add(1)
	factor.Add(2)
	factor.Add(3)

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
	factor.Add(1)

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

	factor.Add(1) // x
	factor.Add(2)
	factor.Add(3) // x
	factor.Add(4)
	factor.Add(5)
	factor.Add(6)
	factor.Add(7) // x

	val, ok := factor.Remove(2)
	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 2)
	}

	if val != 3 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 2, 3)
	}

	if factor.Len() != 6 {
		t.Fatalf("unexpected length %d of array after one removal instead of %d", factor.Len(), 6)
	}

	val, ok = factor.Remove(0)

	if !ok {
		t.Fatalf("unexpected not ok removing by index %d", 0)
	}

	if val != 1 {
		t.Fatalf("unexpected value %d of removed item by index %d instread of %d", val, 0, 1)
	}

	if factor.Len() != 5 {
		t.Fatalf("unexpected length %d of array after two removals instead of %d", factor.Len(), 5)
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
}
