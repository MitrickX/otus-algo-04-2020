package queue

import (
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/homeworks/4/container/list"
)

func TestBinarySearchIn0Array(t *testing.T) {
	array := list.NewArray()

	index, found := BinarySearchInt(array, 100)

	if found {
		t.Errorf("unexpected found in empty array")
	}

	if index != 0 {
		t.Errorf("unexpected index = %d instread of %d", index, 0)
	}
}

func TestBinarySearchIn1ArrayMissInBeginning(t *testing.T) {
	array := list.NewArray()
	array.Append(1)

	index, found := BinarySearchInt(array, 100)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 1 {
		t.Errorf("unexpected index = %d instread of %d", index, 1)
	}
}

func TestBinarySearchIn1ArrayMissInEnding(t *testing.T) {
	array := list.NewArray()
	array.Append(150)

	index, found := BinarySearchInt(array, 100)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 0 {
		t.Errorf("unexpected index = %d instread of %d", index, 0)
	}
}

func TestBinarySearchIn1ArrayHit(t *testing.T) {
	array := list.NewArray()
	array.Append(100)

	index, found := BinarySearchInt(array, 100)

	if !found {
		t.Errorf("unexpected not found")
	}

	if index != 0 {
		t.Errorf("unexpected index = %d instread of %d", index, 0)
	}
}

func TestBinarySearchIn2ArrayMiss0(t *testing.T) {
	array := list.NewArray()
	array.Append(4)
	array.Append(5)

	index, found := BinarySearchInt(array, 3)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 0 {
		t.Errorf("unexpected index = %d instread of %d", index, 0)
	}
}

func TestBinarySearchIn2ArrayMiss1(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(5)

	index, found := BinarySearchInt(array, 3)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 1 {
		t.Errorf("unexpected index = %d instread of %d", index, 1)
	}
}

func TestBinarySearchIn2ArrayMissIn2(t *testing.T) {
	array := list.NewArray()
	array.Append(4)
	array.Append(5)

	index, found := BinarySearchInt(array, 6)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 2 {
		t.Errorf("unexpected index = %d instread of %d", index, 2)
	}
}

func TestBinarySearchIn2ArrayHit0(t *testing.T) {
	array := list.NewArray()
	array.Append(4)
	array.Append(5)

	index, found := BinarySearchInt(array, 4)

	if !found {
		t.Errorf("unexpected not found")
	}

	if index != 0 {
		t.Errorf("unexpected index = %d instread of %d", index, 0)
	}
}

func TestBinarySearchIn2ArrayHit1(t *testing.T) {
	array := list.NewArray()
	array.Append(4)
	array.Append(5)

	index, found := BinarySearchInt(array, 5)

	if !found {
		t.Errorf("unexpected not found")
	}

	if index != 1 {
		t.Errorf("unexpected index = %d instread of %d", index, 1)
	}
}

func TestBinarySearchIn3ArrayMiss0(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(2)
	array.Append(5)

	index, found := BinarySearchInt(array, 0)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 0 {
		t.Errorf("unexpected index = %d instread of %d", index, 0)
	}
}

func TestBinarySearchIn3ArrayMiss1(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(3)
	array.Append(5)

	index, found := BinarySearchInt(array, 2)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 1 {
		t.Errorf("unexpected index = %d instread of %d", index, 1)
	}
}

func TestBinarySearchIn3ArrayMiss2(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(3)
	array.Append(5)

	index, found := BinarySearchInt(array, 4)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 2 {
		t.Errorf("unexpected index = %d instread of %d", index, 2)
	}
}

func TestBinarySearchIn3ArrayMiss3(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(3)
	array.Append(5)

	index, found := BinarySearchInt(array, 6)

	if found {
		t.Errorf("unexpected found")
	}

	if index != 3 {
		t.Errorf("unexpected index = %d instread of %d", index, 3)
	}
}

func TestBinarySearchIn3ArrayHit0(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(2)
	array.Append(5)

	index, found := BinarySearchInt(array, 1)

	if !found {
		t.Errorf("unexpected not found")
	}

	if index != 0 {
		t.Errorf("unexpected index = %d instread of %d", index, 0)
	}
}

func TestBinarySearchIn3ArrayHit1(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(2)
	array.Append(5)

	index, found := BinarySearchInt(array, 2)

	if !found {
		t.Errorf("unexpected not found")
	}

	if index != 1 {
		t.Errorf("unexpected index = %d instread of %d", index, 1)
	}
}

func TestBinarySearchIn3ArrayHit2(t *testing.T) {
	array := list.NewArray()
	array.Append(1)
	array.Append(2)
	array.Append(5)

	index, found := BinarySearchInt(array, 5)

	if !found {
		t.Errorf("unexpected not found")
	}

	if index != 2 {
		t.Errorf("unexpected index = %d instread of %d", index, 2)
	}
}
