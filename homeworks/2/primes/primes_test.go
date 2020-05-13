// +build !actual

package primes

import (
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

// TESTS

func TestCountBruteForce(t *testing.T) {
	testPrimes(t, CountBruteForce, "CountBruteForce", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 10000
	})
}

func TestCountBruteForceWithBreak(t *testing.T) {
	testPrimes(t, CountBruteForceWithBreak, "CountBruteForceWithBreak", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 100000
	})
}

func TestCountSqrtN(t *testing.T) {
	testPrimes(t, CountSqrtN, "CountSqrtN", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 1000000
	})
}

func TestCountSqrtNWithBreak(t *testing.T) {
	testPrimes(t, CountSqrtNWithBreak, "CountSqrtNWithBreak", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 1000000
	})
}

func TestCountSieve(t *testing.T) {
	testPrimes(t, CountSieve, "CountSieve", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 100000000
	})
}

func TestCountSieveWithMemoryOptimization(t *testing.T) {
	testPrimes(t, CountSieveWithMemoryOptimization, "CountSieveWithMemoryOptimization", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 100000000
	})
}

func TestCountFastSieve(t *testing.T) {
	testPrimes(t, CountFastSieve, "CountFastSieve", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 100000000
	})
}

// BENCHMARKS

// n=10^3

func BenchmarkCountBruteForce1000(b *testing.B) {
	n := uint(1000)

	for i := 0; i < b.N; i++ {
		CountBruteForce(n)
	}
}

func BenchmarkCountBruteForceWithBreak1000(b *testing.B) {
	n := uint(1000)

	for i := 0; i < b.N; i++ {
		CountBruteForceWithBreak(n)
	}
}

func BenchmarkCountSqrtN1000(b *testing.B) {
	n := uint(1000)

	for i := 0; i < b.N; i++ {
		CountSqrtN(n)
	}
}

func BenchmarkCountSqrtNWithBreak1000(b *testing.B) {
	n := uint(1000)

	for i := 0; i < b.N; i++ {
		CountSqrtNWithBreak(n)
	}
}

func BenchmarkCountSieve1000(b *testing.B) {
	n := uint(1000)

	for i := 0; i < b.N; i++ {
		CountSieve(n)
	}
}

func BenchmarkCountSieveWithMemoryOptimization1000(b *testing.B) {
	n := uint(1000)

	for i := 0; i < b.N; i++ {
		CountSieveWithMemoryOptimization(n)
	}
}

func BenchmarkCountFastSieve1000(b *testing.B) {
	n := uint(1000)

	for i := 0; i < b.N; i++ {
		CountFastSieve(n)
	}
}

// n=10^4

func BenchmarkCountBruteForce10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		CountBruteForce(n)
	}
}

func BenchmarkCountBruteForceWithBreak10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		CountBruteForceWithBreak(n)
	}
}

func BenchmarkCountSqrtN10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		CountSqrtN(n)
	}
}

func BenchmarkCountSqrtNWithBreak10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		CountSqrtNWithBreak(n)
	}
}

func BenchmarkCountSieve10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		CountSieve(n)
	}
}

func BenchmarkCountSieveWithMemoryOptimization10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		CountSieveWithMemoryOptimization(n)
	}
}

func BenchmarkCountFastSieve10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		CountFastSieve(n)
	}
}

// n=10^5

func BenchmarkCountBruteForceWithBreak100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		CountBruteForceWithBreak(n)
	}
}

func BenchmarkCountSqrtN100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		CountSqrtN(n)
	}
}

func BenchmarkCountSqrtNWithBreak100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		CountSqrtNWithBreak(n)
	}
}

func BenchmarkCountSieve100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		CountSieve(n)
	}
}

func BenchmarkCountSieveWithMemoryOptimization100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		CountSieveWithMemoryOptimization(n)
	}
}

func BenchmarkCountFastSieve100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		CountFastSieve(n)
	}
}

// n=10^6

func BenchmarkCountSqrtN1000000(b *testing.B) {
	n := uint(1000000)

	for i := 0; i < b.N; i++ {
		CountSqrtN(n)
	}
}

func BenchmarkCountSqrtNWithBreak1000000(b *testing.B) {
	n := uint(1000000)

	for i := 0; i < b.N; i++ {
		CountSqrtNWithBreak(n)
	}
}

func BenchmarkCountSieve1000000(b *testing.B) {
	n := uint(1000000)

	for i := 0; i < b.N; i++ {
		CountSieve(n)
	}
}

func BenchmarkCountSieveWithMemoryOptimization1000000(b *testing.B) {
	n := uint(1000000)

	for i := 0; i < b.N; i++ {
		CountSieveWithMemoryOptimization(n)
	}
}

func BenchmarkCountFastSieve1000000(b *testing.B) {
	n := uint(1000000)

	for i := 0; i < b.N; i++ {
		CountFastSieve(n)
	}
}

// n=10^7

func BenchmarkCountSieve10000000(b *testing.B) {
	n := uint(10000000)

	for i := 0; i < b.N; i++ {
		CountSieve(n)
	}
}

func BenchmarkCountSieveWithMemoryOptimization10000000(b *testing.B) {
	n := uint(10000000)

	for i := 0; i < b.N; i++ {
		CountSieveWithMemoryOptimization(n)
	}
}

func BenchmarkCountFastSieve10000000(b *testing.B) {
	n := uint(10000000)

	for i := 0; i < b.N; i++ {
		CountFastSieve(n)
	}
}

// n=10^8

func BenchmarkCountSieve100000000(b *testing.B) {
	n := uint(100000000)

	for i := 0; i < b.N; i++ {
		CountSieve(n)
	}
}

func BenchmarkCountSieveWithMemoryOptimization100000000(b *testing.B) {
	n := uint(100000000)

	for i := 0; i < b.N; i++ {
		CountSieveWithMemoryOptimization(n)
	}
}

func BenchmarkCountFastSieve100000000(b *testing.B) {
	n := uint(100000000)

	for i := 0; i < b.N; i++ {
		CountFastSieve(n)
	}
}

func getCurrentPath() string {
	//nolint:dogsled
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filename))

	return path
}

func getInputsDir() string {
	currentPath := getCurrentPath()
	return filepath.Join(currentPath, "data")
}

func convertToTask(fn func(uint) uint) func(string) string {
	return func(inputData string) string {
		n, _ := strconv.Atoi(inputData)
		result := fn(uint(n))

		return strconv.FormatUint(uint64(result), 10)
	}
}

func testPrimes(t *testing.T, fn func(uint) uint, name string, skip func(string) bool) {
	task := convertToTask(fn)
	taskTester := tester.NewTaskTesterFn(task)
	resultList := taskTester.RunDirWithSkipped(getInputsDir(), skip)

	for _, result := range resultList {
		// test intentionally skipped
		if result.Skipped {
			continue
		}

		// test not run for some reason
		if !result.Run {
			t.Errorf("%s - test #%d has not been run because of %s", name, result.ID, result.Err)
			continue
		}

		if !result.Ok {
			t.Errorf("%s - test #%d is fail, unexpected `%s` instread of `%s`", name, result.ID, result.Actual, result.Expected)
			break
		}
	}
}
