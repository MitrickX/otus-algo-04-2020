// +build !actual

package fibonacci

import (
	"math/big"
	"testing"
)

func TestMultiplyMatrixUint1(t *testing.T) {
	m := Matrix2x2Uint{0, 1, 2, 3}
	w := Matrix2x2Uint{4, 5, 6, 7}

	m.Multiply(w)

	expected := Matrix2x2Uint{6, 7, 26, 31}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMultiplyMatrixUint2(t *testing.T) {
	m := Matrix2x2Uint{0, 1, 2, 3}
	w := Matrix2x2Uint{0, 1, 2, 3}

	m.Multiply(w)

	expected := Matrix2x2Uint{2, 3, 6, 11}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrix2x2UintSymmetricPow2(t *testing.T) {
	m := Matrix2x2Uint{1, 1, 1, 2}
	m.Pow2()

	expected := Matrix2x2Uint{2, 3, 3, 5}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrixUintPow0(t *testing.T) {
	m := Matrix2x2Uint{0, 1, 2, 3}

	m.Pow(0)

	expected := Matrix2x2Uint{1, 0, 0, 1}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrixUintPow1(t *testing.T) {
	m := Matrix2x2Uint{0, 1, 2, 3}

	m.Pow(1)

	expected := Matrix2x2Uint{0, 1, 2, 3}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrixUintPow2(t *testing.T) {
	m := Matrix2x2Uint{0, 1, 2, 3}

	m.Pow(2)

	expected := Matrix2x2Uint{2, 3, 6, 11}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrixUintPow3(t *testing.T) {
	m := Matrix2x2Uint{0, 1, 2, 3}

	m.Pow(3)

	expected := Matrix2x2Uint{6, 11, 22, 39}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrixUintPow4(t *testing.T) {
	m := Matrix2x2Uint{0, 1, 2, 3}

	m.Pow(4)

	expected := Matrix2x2Uint{22, 39, 78, 139}

	if m != expected {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMultiplyMatrix2x2BigInt1(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	w := Matrix2x2BigInt{
		big.NewInt(4), big.NewInt(5),
		big.NewInt(6), big.NewInt(7),
	}

	m.Multiply(w)

	expected := Matrix2x2BigInt{
		big.NewInt(6), big.NewInt(7),
		big.NewInt(26), big.NewInt(31),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMultiplyMatrix2x2BigInt2(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	w := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	m.Multiply(w)

	expected := Matrix2x2BigInt{
		big.NewInt(2), big.NewInt(3),
		big.NewInt(6), big.NewInt(11),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMultiplyMatrix2x2BigInt3(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	m.Multiply(m)

	expected := Matrix2x2BigInt{
		big.NewInt(2), big.NewInt(3),
		big.NewInt(6), big.NewInt(11),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrix2x2BigIntSymmetricPow2(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(1), big.NewInt(1),
		big.NewInt(1), big.NewInt(2),
	}
	m.Pow2()

	expected := Matrix2x2BigInt{
		big.NewInt(2), big.NewInt(3),
		big.NewInt(3), big.NewInt(5),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrix2x2BigIntPow0(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}
	m.Pow(0)

	expected := Matrix2x2BigInt{
		big.NewInt(1), big.NewInt(0),
		big.NewInt(0), big.NewInt(1),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrix2x2BigIntPow1(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	m.Pow(1)

	expected := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrix2x2BigIntPow2(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	m.Pow(2)

	expected := Matrix2x2BigInt{
		big.NewInt(2), big.NewInt(3),
		big.NewInt(6), big.NewInt(11),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrix2x2BigIntPow3(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	m.Pow(3)

	expected := Matrix2x2BigInt{
		big.NewInt(6), big.NewInt(11),
		big.NewInt(22), big.NewInt(39),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}

func TestMatrix2x2BigIntPow4(t *testing.T) {
	m := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(2), big.NewInt(3),
	}

	m.Pow(4)

	expected := Matrix2x2BigInt{
		big.NewInt(22), big.NewInt(39),
		big.NewInt(78), big.NewInt(139),
	}

	if !m.Equals(&expected) {
		t.Errorf("unexpected %v instead of %v", m, expected)
	}
}
