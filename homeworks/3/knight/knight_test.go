package knight

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

func TestWalks(t *testing.T) {
	taskTester := tester.NewTaskTesterFn(func(inputData string) string {
		x, _ := strconv.Atoi(inputData)
		count, mask := Walks(x)
		return fmt.Sprintf("%d\n%d", count, mask)
	})

	inputsDir := getInputsDir()
	resultList := taskTester.RunDir(inputsDir)

	for _, result := range resultList {
		// test intentionally skipped
		if result.Skipped {
			continue
		}

		// test not run for some reason
		if !result.Run {
			t.Errorf("test #%d has not been run because of %s", result.ID, result.Err)
			continue
		}

		if !result.Ok {
			t.Errorf("test #%d is fail, unexpected `%s` instread of `%s`", result.ID, result.Actual, result.Expected)
			break
		}
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
