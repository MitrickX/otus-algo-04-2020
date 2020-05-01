package strlen

import (
	"github.com/mitrickx/otus-algo-04-2020/tester"
	"path/filepath"
	"runtime"
	"testing"
)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filename))
	return path
}

func TestStrLen(t *testing.T) {
	currentPath := getCurrentPath()
	dataDir := filepath.Join(currentPath, "data")

	taskTester := tester.NewTaskTester(StrLen)
	resultList := taskTester.RunDir(dataDir)

	for _, result := range resultList {

		if !result.Run {
			t.Errorf("Test #%d has not been run because of %s", result.Id, result.Err)
			continue
		}

		if !result.Ok {
			t.Errorf("Test #%d is fail", result.Id)
		}
	}
}