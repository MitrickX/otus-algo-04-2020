package fibonacci

import (
	"math"
	"math/big"
)

var (
	// sqrt5 precomputed
	sqrt5 float64 = math.Sqrt(5)
	// sqrt5Big precomputed
	sqrt5Big *big.Float = big.NewFloat(sqrt5)
	// phiBig precomputed
	phiBig *big.Float = big.NewFloat(math.Phi)
)

func Recursion(n uint) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	return Recursion(n-1) + Recursion(n-2)
}

func Iteration(n uint) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	prev := uint64(0)
	cur := uint64(1)

	for i := uint(2); i <= n; i++ {
		next := prev + cur
		prev = cur
		cur = next
	}

	return cur
}

func GoldenRatio(n uint) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	phiIntoN := math.Pow(math.Phi, float64(n))
	fraction := math.Round(phiIntoN / sqrt5)

	return uint64(fraction)
}

func Matrix(n uint) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	p := Matrix2x2Uint{0, 1, 1, 1}
	p.Pow(n - 1)

	return p[3]
}

func IterationBigInt(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	prev := big.NewInt(0)
	cur := big.NewInt(1)

	// allocate onetime
	next := new(big.Int)

	for i := uint(2); i <= n; i++ {
		next.Add(prev, cur) // next = prev + cur
		prev.Set(cur)       // prev = cur
		cur.Set(next)       // cur = next
	}

	return cur
}

func GoldenRatioBigInt(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	phiIntoN := new(big.Float)
	phiIntoN.Copy(phiBig)
	Pow(phiIntoN, n)

	// divide
	div := new(big.Float)
	div.Quo(phiIntoN, sqrt5Big)
	div.Add(div, big.NewFloat(0.5))

	result, _ := div.Int(nil)

	return result
}

func MatrixBigInt(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	p := Matrix2x2BigInt{
		big.NewInt(0), big.NewInt(1),
		big.NewInt(1), big.NewInt(1),
	}
	p.Pow(n - 1)

	return p[3]
}
