package fibonacci

import (
	"math/big"
)

func Pow(x *big.Float, n uint) {
	if n == 0 {
		x.SetFloat64(1)
		return
	}

	if n == 1 {
		return
	}

	y := new(big.Float)
	y.Copy(x)

	x.SetFloat64(1)

	for n > 0 {
		// odd case
		if n&1 == 1 {
			n--

			// a^n = a * a^(n-1)
			x.Mul(x, y)
		}

		// n now is even, divide by 2 for next step
		n >>= 1

		// a^n = a^(n/2) * a^(n/2)
		y.Mul(y, y)
	}
}
