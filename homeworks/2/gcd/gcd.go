package gcd

import "math/big"

// zero is pre calculated big int 0.
var zero = big.NewInt(0)

// Subtract calculate GCD by naive subtract algorithm.
func Subtract(x *big.Int, y *big.Int) *big.Int {
	// avoid side-effect - numbers changed for outer word - clone numbers
	a := new(big.Int).Set(x)
	b := new(big.Int).Set(y)

	for a.Cmp(b) != 0 {
		if a.Cmp(b) > 0 {
			a.Sub(a, b)
		} else {
			b.Sub(b, a)
		}
	}

	return a
}

// Rest calculate GCD by find rest algorithm (modulo).
func Rest(x *big.Int, y *big.Int) *big.Int {
	// avoid side-effect - numbers changed for outer word - clone numbers
	a := new(big.Int).Set(x)
	b := new(big.Int).Set(y)

	for a.Cmp(zero) > 0 && b.Cmp(zero) > 0 {
		if a.Cmp(b) > 0 {
			a.Mod(a, b) // a = a%b
		} else {
			b.Mod(b, a) // b = b%a
		}
	}

	if a.Cmp(zero) > 0 {
		return a
	}

	return b
}

// Binary calculate GCD by Stein's algorithm.
func Binary(x *big.Int, y *big.Int) *big.Int {
	// avoid side-effect - numbers changed for outer word - clone numbers
	a := new(big.Int).Set(x)
	b := new(big.Int).Set(y)

	// it is buffer number to prevent redundant allocation of memory
	c := new(big.Int)

	return binary(a, b, c)
}

// binary calculate GCD by Stein's algorithm, c here is just for optimization to prevent redundant allocation of memory.
func binary(a *big.Int, b *big.Int, c *big.Int) *big.Int {
	if a.Cmp(b) == 0 {
		return a
	}

	if a.Cmp(zero) == 0 {
		return b
	}

	if b.Cmp(zero) == 0 {
		return a
	}

	// a is even
	if a.Bit(0) == 0 {
		// b is odd
		if b.Bit(0) == 1 {
			return binary(a.Rsh(a, 1), b, c)
		}

		// a and b are both even
		c = binary(a.Rsh(a, 1), b.Rsh(b, 1), c)

		return c.Lsh(c, 1)
	}

	// b is event and a is odd
	if b.Bit(0) == 0 {
		return binary(a, b.Rsh(b, 1), c)
	}

	if a.Cmp(b) > 0 {
		c.Sub(a, b)
		c.Rsh(c, 1)

		return binary(c, b, a)
	}

	c.Sub(b, a)
	c.Rsh(c, 1)

	return binary(c, a, b)
}
