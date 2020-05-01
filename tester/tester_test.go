package tester

import (
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

var testTask = Task(func(inputData string) string {
	n, _ := strconv.Atoi(inputData)
	return strconv.Itoa(n * n)
})

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filename))
	return path
}

func TestTaskTester_RunDir(t *testing.T) {
	tester := NewTaskTester(testTask)

	currentPath := getCurrentPath()
	dataDir := filepath.Join(currentPath, "data")
	resultList := tester.RunDir(dataDir)

	for _, result := range resultList {

		// 3d test is intentionally broken - expected by fail
		if result.Id == 3 {
			if !result.Run {
				t.Errorf("Test #%d has not been run because of %s", result.Id, result.Err)
			} else if result.Ok {
				t.Errorf("Test #%d expected be fail", result.Id)
			}
			continue
		}

		// 4th test can't be run, cause of output file is not exists
		if result.Id == 4 {
			if result.Run {
				t.Errorf("Test #%d expected to be not run", result.Id)
			}
			if result.Err != TestErrorOutputFileNotFound {
				t.Errorf("Test #%d expected to be has error `%s` instread of `%s`",
					result.Id,
					TestErrorOutputFileNotFound,
					result.Err)
			}
			continue
		}

		// 5th test can't be run, cause of input file is not exists
		if result.Id == 5 {
			if result.Run {
				t.Errorf("Test #%d expected to be not run", result.Id)
			}
			if result.Err != TestErrorInputFileNotFound {
				t.Errorf("Test #%d expected to be has error `%s` instread of `%s`",
					result.Id,
					TestErrorInputFileNotFound,
					result.Err)
			}
			continue
		}

		// Other tests are good
		if !result.Run {
			t.Errorf("Test #%d has not been run because of %s", result.Id, result.Err)
		} else if !result.Ok {
			t.Errorf("Test #%d is fail", result.Id)
		}

	}

}
