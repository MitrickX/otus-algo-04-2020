package primes

import (
	"math"
)

// CountPrimes count how many is prime numbers from (1..n] by isPrime function.
func CountPrimes(n uint, isPrime func(uint) bool) uint {
	count := uint(0)

	for i := uint(2); i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

// CountBruteForce count how many is prime numbers from (1..n] by IsPrimeBruteForce.
func CountBruteForce(n uint) uint {
	return CountPrimes(n, IsPrimeBruteForce)
}

// CountBruteForceWithBreak count how many is prime numbers from (1..n] by IsPrimeBruteForceWithBreak.
func CountBruteForceWithBreak(n uint) uint {
	return CountPrimes(n, IsPrimeBruteForceWithBreak)
}

// CountSqrtN count how many is prime numbers from (1..n] by IsPrimeSqrtN.
func CountSqrtN(n uint) uint {
	return CountPrimes(n, IsPrimeSqrtN)
}

// CountSqrtN count how many is prime numbers from (1..n] by IsPrimeSqrtNWithBreak.
func CountSqrtNWithBreak(n uint) uint {
	return CountPrimes(n, IsPrimeSqrtNWithBreak)
}

// CountSieve count how many is prime numbers from (1..n] by Eratosthenes Sieve.
func CountSieve(n uint) uint {
	if n <= 1 {
		return 0
	}

	// sieve memory
	sieve := make([]bool, n+1)

	sqrtN := uint(math.Sqrt(float64(n)) + 1)

	// sieving loop
	for i := uint(2); i <= sqrtN; i++ {
		if sieve[i] {
			continue
		}

		// optimization, step for inner loop == i only if i == 2
		step := 2 * i
		if i == 2 {
			step /= 2
		}

		for j := i * i; j <= n; j += step {
			sieve[j] = true
		}
	}

	// count not sieved - this is primes number
	count := uint(0)

	for i := uint(2); i <= n; i++ {
		if !sieve[i] {
			count++
		}
	}

	return count
}

// CountSieveWithMemoryOptimization count how many is prime numbers from (1..n] by Eratosthenes Sieve,
// where Sieve is matrix bit implementation.
func CountSieveWithMemoryOptimization(n uint) uint {
	if n <= 1 {
		return 0
	}

	// sieve memory, bit matrix
	sieve := NewSieve(n)

	sqrtN := uint(math.Sqrt(float64(n)) + 1)

	// sieving loop
	for i := uint(2); i <= sqrtN; i++ {
		if sieve.IsSieved(i) {
			continue
		}

		// optimization, step for inner loop == i only if i == 2
		step := 2 * i
		if i == 2 {
			step /= 2
		}

		for j := i * i; j <= n; j += step {
			sieve.SetSieved(j)
		}
	}

	// count not sieved - this is primes number
	count := uint(0)

	for i := uint(2); i <= n; i++ {
		if !sieve.IsSieved(i) {
			count++
		}
	}

	return count
}

// CountFastSieve count how many is prime numbers from (1..n] by Eratosthenes Sieve with O(n) complexity
// But this method eat a lot of memory, so not so optimal on practice.
func CountFastSieve(n uint) uint {
	if n <= 1 {
		return 0
	}

	// least primes map
	lp := make([]uint, n+1)

	// primes list
	primes := []uint(nil)

	for i := uint(2); i <= n; i++ {
		// if least primes for current i is 0 than i is prime
		if lp[i] == 0 {
			lp[i] = i
			primes = append(primes, i)
		}

		// fill least primes map for all x = p * i for p <= lp[i]
		for _, p := range primes {
			x := p * i

			canFill := p <= lp[i] && x <= n
			if !canFill {
				break
			}

			lp[x] = p
		}
	}

	return uint(len(primes))
}

// isPrimeBruteForce checks is n prime number by the most inefficient way.
func IsPrimeBruteForce(n uint) bool {
	count := 0

	for i := uint(2); i < n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count == 0
}

// isPrimeBruteForceWithCut checks is n prime number by brute force but with break optimization.
func IsPrimeBruteForceWithBreak(n uint) bool {
	for i := uint(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// IsPrimeSqrtN checks is n prime number by brute force to the sqrt of N.
func IsPrimeSqrtN(n uint) bool {
	count := 0
	sqrtN := uint(math.Sqrt(float64(n)))

	for i := uint(2); i <= sqrtN; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count == 0
}

// IsPrimeSqrtN checks is n prime number by brute force to the sqrt of N with break optimization.
func IsPrimeSqrtNWithBreak(n uint) bool {
	sqrtN := uint(math.Sqrt(float64(n)))

	for i := uint(2); i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
