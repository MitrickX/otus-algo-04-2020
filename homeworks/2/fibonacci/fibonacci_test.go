package fibonacci

import (
	"math/big"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

// TESTS

func TestRecursionUint64(t *testing.T) {
	testFibonacciUint64(t, Recursion, "Recursion", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n > 40
	})
}

func TestIterationUint64(t *testing.T) {
	testFibonacciUint64(t, Iteration, "Iteration", nil)
}

func TestGoldenRatioUint64(t *testing.T) {
	testFibonacciUint64(t, Iteration, "GoldenRatio", nil)
}

func TestMatrixUint64(t *testing.T) {
	testFibonacciUint64(t, Matrix, "Matrix", nil)
}

func TestIterationBigInt(t *testing.T) {
	testFibonacciBigInt64(t, IterationBigInt, "IterationBigInt", nil)
}

func TestGoldenRatioBigInt(t *testing.T) {
	testFibonacciBigInt64(t, GoldenRatioBigInt, "GoldenRatioBigInt", func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n >= 93
	})
}

func TestMatrixBigInt(t *testing.T) {
	testFibonacciBigInt64(t, MatrixBigInt, "MatrixBigInt", nil)
}

// BENCHMARKS

func BenchmarkRecursion3(b *testing.B) {
	n := uint(3)

	for i := 0; i < b.N; i++ {
		_ = Recursion(n)
	}
}

func BenchmarkIteration3(b *testing.B) {
	n := uint(3)

	for i := 0; i < b.N; i++ {
		_ = Iteration(n)
	}
}

func BenchmarkGoldenRatio3(b *testing.B) {
	n := uint(3)

	for i := 0; i < b.N; i++ {
		_ = GoldenRatio(n)
	}
}

func BenchmarkMatrix3(b *testing.B) {
	n := uint(3)

	for i := 0; i < b.N; i++ {
		_ = Matrix(n)
	}
}

func BenchmarkRecursion7(b *testing.B) {
	n := uint(7)

	for i := 0; i < b.N; i++ {
		_ = Recursion(n)
	}
}

func BenchmarkIteration7(b *testing.B) {
	n := uint(7)

	for i := 0; i < b.N; i++ {
		_ = Iteration(n)
	}
}

func BenchmarkGoldenRatio7(b *testing.B) {
	n := uint(7)

	for i := 0; i < b.N; i++ {
		_ = GoldenRatio(n)
	}
}

func BenchmarkMatrix7(b *testing.B) {
	n := uint(7)

	for i := 0; i < b.N; i++ {
		_ = Matrix(n)
	}
}

func BenchmarkRecursion15(b *testing.B) {
	n := uint(15)

	for i := 0; i < b.N; i++ {
		_ = Recursion(n)
	}
}

func BenchmarkIteration15(b *testing.B) {
	n := uint(15)

	for i := 0; i < b.N; i++ {
		_ = Iteration(n)
	}
}

func BenchmarkGoldenRatio15(b *testing.B) {
	n := uint(15)

	for i := 0; i < b.N; i++ {
		_ = GoldenRatio(n)
	}
}

func BenchmarkMatrix15(b *testing.B) {
	n := uint(15)

	for i := 0; i < b.N; i++ {
		_ = Matrix(n)
	}
}

func BenchmarkRecursion20(b *testing.B) {
	n := uint(20)

	for i := 0; i < b.N; i++ {
		_ = Recursion(n)
	}
}

func BenchmarkIteration20(b *testing.B) {
	n := uint(20)

	for i := 0; i < b.N; i++ {
		_ = Iteration(n)
	}
}

func BenchmarkGoldenRatio20(b *testing.B) {
	n := uint(20)

	for i := 0; i < b.N; i++ {
		_ = GoldenRatio(n)
	}
}

func BenchmarkMatrix20(b *testing.B) {
	n := uint(20)

	for i := 0; i < b.N; i++ {
		_ = Matrix(n)
	}
}

func BenchmarkRecursion40(b *testing.B) {
	n := uint(40)

	for i := 0; i < b.N; i++ {
		_ = Recursion(n)
	}
}

func BenchmarkIteration40(b *testing.B) {
	n := uint(40)

	for i := 0; i < b.N; i++ {
		_ = Iteration(n)
	}
}

func BenchmarkGoldenRatio40(b *testing.B) {
	n := uint(40)

	for i := 0; i < b.N; i++ {
		_ = GoldenRatio(n)
	}
}

func BenchmarkMatrix40(b *testing.B) {
	n := uint(40)

	for i := 0; i < b.N; i++ {
		_ = Matrix(n)
	}
}

func BenchmarkIteration93(b *testing.B) {
	n := uint(93)

	for i := 0; i < b.N; i++ {
		_ = Iteration(n)
	}
}

func BenchmarkGoldenRatio93(b *testing.B) {
	n := uint(93)

	for i := 0; i < b.N; i++ {
		_ = GoldenRatio(n)
	}
}

func BenchmarkMatrix93(b *testing.B) {
	n := uint(93)

	for i := 0; i < b.N; i++ {
		_ = Matrix(n)
	}
}

func BenchmarkIteration184(b *testing.B) {
	n := uint(184)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix184(b *testing.B) {
	n := uint(184)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func BenchmarkIteration300(b *testing.B) {
	n := uint(300)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix300(b *testing.B) {
	n := uint(300)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func BenchmarkIteration501(b *testing.B) {
	n := uint(501)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix501(b *testing.B) {
	n := uint(501)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func BenchmarkIteration999(b *testing.B) {
	n := uint(999)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix999(b *testing.B) {
	n := uint(999)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func BenchmarkIteration1001(b *testing.B) {
	n := uint(1001)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix1001(b *testing.B) {
	n := uint(1001)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func getCurrentPath() string {
	//nolint:dogsled
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filename))

	return path
}

func getIntDataDir() string {
	currentPath := getCurrentPath()
	return filepath.Join(currentPath, "data", "int")
}

func getBigIntDataDir() string {
	currentPath := getCurrentPath()
	return filepath.Join(currentPath, "data", "bigint")
}

func convertToTaskUint64(fn func(uint) uint64) func(string) string {
	return func(inputData string) string {
		n, _ := strconv.Atoi(inputData)
		result := fn(uint(n))

		return strconv.FormatUint(result, 10)
	}
}

func convertToTaskBigInt(fn func(uint) *big.Int) func(string) string {
	return func(inputData string) string {
		n, _ := strconv.Atoi(inputData)
		result := fn(uint(n))

		return result.String()
	}
}

func testFibonacciUint64(t *testing.T, fn func(uint) uint64, name string, skip func(string) bool) {
	task := convertToTaskUint64(fn)
	taskTester := tester.NewTaskTesterFn(task)
	resultList := taskTester.RunDirWithSkipped(getIntDataDir(), skip)

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

func testFibonacciBigInt64(t *testing.T, fn func(uint) *big.Int, name string, skip func(string) bool) {
	task := convertToTaskBigInt(fn)
	taskTester := tester.NewTaskTesterFn(task)
	resultList := taskTester.RunDirWithSkipped(getBigIntDataDir(), skip)

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
