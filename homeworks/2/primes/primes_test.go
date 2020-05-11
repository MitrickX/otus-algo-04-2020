package primes

import (
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

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
