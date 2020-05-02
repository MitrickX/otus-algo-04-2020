package tester

import (
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

var testTask = TaskFn(func(inputData string) string {
	n, _ := strconv.Atoi(inputData)
	return strconv.Itoa(n * n)
})

func getCurrentPath() string {
	//nolint:dogsled
	//nolint:gomnd
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Dir(filename))

	return path
}

func TestTaskTester_RunDir(t *testing.T) {
	tester := NewTaskTesterFn(testTask)

	currentPath := getCurrentPath()
	dataDir := filepath.Join(currentPath, "data")
	resultList := tester.RunDir(dataDir)

	// 3d test is intentionally broken - expected by fail
	idBroken := 3

	// 4th test can't be run, cause of output file is not exists
	idOutputNotFound := 4

	// 5th test can't be run, cause of input file is not exists
	idInputNotFound := 5

	for _, result := range resultList {
		if result.ID == idBroken {
			if !result.Run {
				t.Errorf("Test #%d has not been run because of %s", result.ID, result.Err)
			} else if result.Ok {
				t.Errorf("Test #%d expected be fail", result.ID)
			}

			continue
		}

		if result.ID == idOutputNotFound {
			if result.Run {
				t.Errorf("Test #%d expected to be not run", result.ID)
			} else if result.Err != ErrOutputFileNotFound {
				t.Errorf("Test #%d expected to be has error `%s` instread of `%s`",
					result.ID,
					ErrOutputFileNotFound,
					result.Err)
			}

			continue
		}

		if result.ID == idInputNotFound {
			if result.Run {
				t.Errorf("Test #%d expected to be not run", result.ID)
			} else if result.Err != ErrInputFileNotFound {
				t.Errorf("Test #%d expected to be has error `%s` instread of `%s`",
					result.ID,
					ErrInputFileNotFound,
					result.Err)
			}

			continue
		}

		// Other tests are good
		if !result.Run {
			t.Errorf("Test #%d has not been run because of %s", result.ID, result.Err)
		} else if !result.Ok {
			t.Errorf("Test #%d is fail", result.ID)
		}
	}
}
