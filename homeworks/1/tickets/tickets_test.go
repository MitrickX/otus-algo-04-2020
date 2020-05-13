// +build !actual

package tickets

import (
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

func getCurrentPath() string {
	//nolint:dogsled
	//nolint:gomnd
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filename))

	return path
}

func getDataDir() string {
	currentPath := getCurrentPath()
	return filepath.Join(currentPath, "data")
}

func TestCalculate(t *testing.T) {
	calculateTester := tester.NewTaskTesterFn(func(inputData string) string {
		n, _ := strconv.Atoi(inputData)
		result := Calculate(n)
		return strconv.FormatInt(result, 10)
	})

	dataDir := getDataDir()
	resultList := calculateTester.RunDir(dataDir)

	for _, result := range resultList {
		if !result.Run {
			t.Errorf("Test #%d has not been run because of %s", result.ID, result.Err)
		} else if !result.Ok {
			t.Errorf("Test #%d is fail, unexpected `%s` instread of `%s`", result.ID, result.Actual, result.Expected)
		}
	}
}
