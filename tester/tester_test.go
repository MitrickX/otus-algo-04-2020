package tester

import (
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

var strLen = Task(func(inputData string) string {
	return strconv.Itoa(len(inputData))
})

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filename))
	return path
}

func TestTaskTester_RunDir(t *testing.T) {
	tester := NewTaskTester(strLen)

	currentPath := getCurrentPath()
	dataDir := filepath.Join(currentPath, "data")
	resultList := tester.RunDir(dataDir)

	for _, result := range resultList {
		if !result.Run {
			t.Errorf("Test #%d has not been run because of %s", result.Id, result.Err)
		} else if !result.Ok {
			t.Errorf("Test #%d is fail", result.Id)
		}
	}

}
