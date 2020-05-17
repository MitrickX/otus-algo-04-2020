package matrix

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/container"
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

	slice := container.Convert(array)
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

	slice := container.Convert(array)
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

	slice := container.Convert(array)
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

	slice := container.Convert(array)
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

	slice := container.Convert(array)
	expected := []interface{}{1, 2, 3, 4, 6, 100, 7, 9, 10}

	if !reflect.DeepEqual(slice, expected) {
		t.Fatalf("unexpected %v instead of %v", slice, expected)
	}
}

// BENCHMARKS

// 10^2.

func BenchmarkAppend100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendN(100)
	}
}

func BenchmarkAdd100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddN(100)
	}
}

func BenchmarkSet100(b *testing.B) {
	ar := AppendN(100)

	for i := 0; i < b.N; i++ {
		SetN(100, ar)
	}
}

func BenchmarkGet100(b *testing.B) {
	ar := AppendN(100)

	for i := 0; i < b.N; i++ {
		GetN(100, ar)
	}
}

func BenchmarkAppendRemove100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendRemoveN(100)
	}
}

func BenchmarkAddRemove100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddRemoveN(100)
	}
}

// 10^3.

func BenchmarkAppend1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendN(1000)
	}
}

func BenchmarkAdd1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddN(1000)
	}
}

func BenchmarkSet1000(b *testing.B) {
	ar := AppendN(1000)

	for i := 0; i < b.N; i++ {
		SetN(1000, ar)
	}
}

func BenchmarkGet1000(b *testing.B) {
	ar := AppendN(1000)

	for i := 0; i < b.N; i++ {
		GetN(1000, ar)
	}
}

func BenchmarkAppendRemove1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendRemoveN(1000)
	}
}

func BenchmarkAddRemove1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddRemoveN(1000)
	}
}

// 10^4.

func BenchmarkAppend10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendN(10000)
	}
}

func BenchmarkAdd10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddN(10000)
	}
}

func BenchmarkSet10000(b *testing.B) {
	ar := AppendN(10000)

	for i := 0; i < b.N; i++ {
		SetN(10000, ar)
	}
}

func BenchmarkGet10000(b *testing.B) {
	ar := AppendN(10000)

	for i := 0; i < b.N; i++ {
		GetN(10000, ar)
	}
}

func BenchmarkAppendRemove10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendRemoveN(10000)
	}
}

func BenchmarkAddRemove10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddRemoveN(10000)
	}
}

// 10^5.

func BenchmarkAppend100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendN(100000)
	}
}

func BenchmarkAdd100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddN(100000)
	}
}

func BenchmarkSet100000(b *testing.B) {
	ar := AppendN(100000)

	for i := 0; i < b.N; i++ {
		SetN(100000, ar)
	}
}

func BenchmarkGet100000(b *testing.B) {
	ar := AppendN(100000)

	for i := 0; i < b.N; i++ {
		GetN(100000, ar)
	}
}

func BenchmarkAppendRemove100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendRemoveN(100000)
	}
}

func BenchmarkAddRemove100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddRemoveN(100000)
	}
}

func AppendN(n int) *Array {
	ar := NewArray()

	for i := 0; i < n; i++ {
		ar.Append(rand.Int()) //nolint:gosec
	}

	return ar
}

func AddN(n int) *Array {
	ar := NewArray()
	ar.Append(rand.Int()) //nolint:gosec

	for i := 0; i < n; i++ {
		ar.Add(rand.Int(), rand.Intn(ar.Len())) //nolint:gosec
	}

	return ar
}

func SetN(n int, ar *Array) {
	for i := 0; i < n; i++ {
		ar.Set(rand.Int(), rand.Intn(ar.Len())) //nolint:gosec
	}
}

func GetN(n int, ar *Array) {
	for i := 0; i < n; i++ {
		ar.Get(rand.Intn(ar.Len())) //nolint:gosec
	}
}

func RemoveAll(ar *Array) {
	for ar.Len() > 0 {
		ar.Remove(rand.Intn(ar.Len())) //nolint:gosec
	}
}

func AppendRemoveN(n int) {
	ar := AppendN(n)
	RemoveAll(ar)
}

func AddRemoveN(n int) {
	ar := AddN(n)
	RemoveAll(ar)
}
