package fibonacci

import (
	"math/big"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

// TESTS FOR SMALL INPUTS

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

// TESTS FOR MIDDLE INPUTS

func TestIterationMiddle(t *testing.T) {
	testFibonacciBigInt(t, IterationBigInt, "TestIterationMiddle", getMiddleInputsDir(), func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n >= 10000000
	})
}

func TestGoldenRatioMiddle(t *testing.T) {
	testFibonacciBigInt(t, GoldenRatioBigInt, "TestGoldenRatioMiddle", getMiddleInputsDir(), func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n >= 93
	})
}

func TestMatrixBigMiddle(t *testing.T) {
	testFibonacciBigInt(t, MatrixBigInt, "TestMatrixBigMiddle", getMiddleInputsDir(), func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n >= 10000000
	})
}

// TESTS FOR BIG INPUTS

func TestIterationBig(t *testing.T) {
	testFibonacciBigInt(t, IterationBigInt, "TestIterationBig", getBigInputsDir(), func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n >= 10000000
	})
}

func TestMatrixBig(t *testing.T) {
	testFibonacciBigInt(t, MatrixBigInt, "TestMatrixBig", getBigInputsDir(), func(inputData string) bool {
		n, _ := strconv.Atoi(inputData)
		return n >= 100000000
	})
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

func BenchmarkIteration10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix10000(b *testing.B) {
	n := uint(10000)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func BenchmarkIteration100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix100000(b *testing.B) {
	n := uint(100000)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func BenchmarkIteration1000000(b *testing.B) {
	n := uint(1000000)

	for i := 0; i < b.N; i++ {
		_ = IterationBigInt(n)
	}
}

func BenchmarkMatrix1000000(b *testing.B) {
	n := uint(1000000)

	for i := 0; i < b.N; i++ {
		_ = MatrixBigInt(n)
	}
}

func BenchmarkMatrix10000000(b *testing.B) {
	n := uint(10000000)

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

func getSmallInputsDir() string {
	currentPath := getCurrentPath()
	return filepath.Join(currentPath, "data", "small")
}

func getMiddleInputsDir() string {
	currentPath := getCurrentPath()
	return filepath.Join(currentPath, "data", "middle")
}

func getBigInputsDir() string {
	currentPath := getCurrentPath()
	return filepath.Join(currentPath, "data", "big")
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
	resultList := taskTester.RunDirWithSkipped(getSmallInputsDir(), skip)

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

func testFibonacciBigInt(t *testing.T, fn func(uint) *big.Int, name string, dir string, skip func(string) bool) {
	task := convertToTaskBigInt(fn)
	taskTester := tester.NewTaskTesterFn(task)
	resultList := taskTester.RunDirWithSkipped(dir, skip)

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
