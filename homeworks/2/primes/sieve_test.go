// +build !actual

package primes

import (
	"testing"
)

func TestNewSieve0(t *testing.T) {
	s := NewSieve(0)
	ln := len(s)

	expectedLen := 0
	if ln != expectedLen {
		t.Errorf("unexpected length %d instread of %d", ln, expectedLen)
	}
}
func TestNewSieve1(t *testing.T) {
	s := NewSieve(1)
	ln := len(s)

	expectedLen := 1
	if ln != expectedLen {
		t.Errorf("unexpected length %d instread of %d", ln, expectedLen)
	}
}

func TestNewSieve31(t *testing.T) {
	s := NewSieve(31)
	ln := len(s)

	expectedLen := 1
	if ln != expectedLen {
		t.Errorf("unexpected length %d instread of %d", ln, expectedLen)
	}
}

func TestNewSieve100(t *testing.T) {
	s := NewSieve(100)
	ln := len(s)

	expectedLen := 4
	if ln != expectedLen {
		t.Errorf("unexpected length %d instread of %d", ln, expectedLen)
	}
}

func TestNewSieve64(t *testing.T) {
	s := NewSieve(64)
	ln := len(s)

	expectedLen := 2
	if ln != expectedLen {
		t.Errorf("unexpected length %d instread of %d", ln, expectedLen)
	}
}

func TestSieve_IsSieved(t *testing.T) {
	s := NewSieve(100)

	ns := []uint{1, 2, 3, 31, 32, 100}
	for _, n := range ns {
		if s.IsSieved(n) {
			t.Errorf("unexpected that new sieve has sieved bit for %d", n)
		}
	}
}

func TestSieve_SetSieved64(t *testing.T) {
	s := NewSieve(64)

	sieved := []uint{2, 3, 31, 32, 64}
	notSieved := []uint{1, 5, 33, 63}

	for _, n := range sieved {
		s.SetSieved(n)
	}

	for _, n := range sieved {
		if !s.IsSieved(n) {
			t.Fatalf("unexpected that not sieved bit for %d", n)
		}
	}

	for _, n := range notSieved {
		if s.IsSieved(n) {
			t.Fatalf("unexpected that sieved bit for %d", n)
		}
	}
}

func TestSieve_SetSieved100(t *testing.T) {
	s := NewSieve(100)

	sieved := []uint{2, 3, 31, 32, 64, 98, 100}
	notSieved := []uint{1, 5, 33, 63, 99}

	for _, n := range sieved {
		s.SetSieved(n)
	}

	for _, n := range sieved {
		if !s.IsSieved(n) {
			t.Fatalf("unexpected that not sieved bit for %d", n)
		}
	}

	for _, n := range notSieved {
		if s.IsSieved(n) {
			t.Fatalf("unexpected that sieved bit for %d", n)
		}
	}
}
