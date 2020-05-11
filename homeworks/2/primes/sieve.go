package primes

import (
	"math"
)

// Sieve for Eratosthenes Sieve algorithm
// Use bit matrix based on uint32
// Sieve supports numbers of interval [1..n], 0 excluded.
type Sieve []uint32

func NewSieve(n uint) Sieve {
	size := uint(math.Ceil(float64(n) / 32))
	return make([]uint32, size)
}

func (s Sieve) IsSieved(k uint) bool {
	if k == 0 {
		return false
	}

	// for 32 bit must be in first integer and shift must be 31, that is why k--
	k--

	// row (what integer in vector, numeration from 0 of of course) in sieve bit matrix
	row := k / 32

	// col in sieve bit matrix, col == k % 32, shift inside integer to check up
	col := k - row*32

	// check is proper bit (by bitwise or)
	return (s[row] & (1 << col)) > 0
}

func (s Sieve) SetSieved(k uint) {
	if k == 0 {
		return
	}

	// for 32 bit must be in first integer and shift must be 31, that is why k--
	k--

	// row (what integer in vector, numeration from 0 of of course) in sieve bit matrix
	row := k / 32

	// col in sieve bit matrix, col == k % 32, shift inside integer to set up
	col := k - row*32

	// set proper bit to 1 (by bitwise or)
	s[row] |= 1 << col
}
