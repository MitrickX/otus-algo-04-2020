package strlen

import (
	"path/filepath"
	"runtime"
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

func TestStrLen(t *testing.T) {
	currentPath := getCurrentPath()
	dataDir := filepath.Join(currentPath, "data")

	taskTester := tester.NewTaskTesterFn(StrLen)
	resultList := taskTester.RunDir(dataDir)

	// 4th test is intentionally broken, rest test are good
	idBroken := 4

	for _, result := range resultList {
		if result.ID != idBroken {
			if !result.Run {
				t.Errorf("Test #%d has not been run because of %s", result.ID, result.Err)
			} else if !result.Ok {
				t.Errorf("Test #%d is fail", result.ID)
			}
		} else {
			if !result.Run {
				t.Errorf("Test #%d has not been run because of %s", result.ID, result.Err)
			} else if result.Ok {
				t.Errorf("Test #%d expected to be fail", result.ID)
			}
		}
	}
}
