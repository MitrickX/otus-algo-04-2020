package tester

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// ErrInputFileNotFound predefined error.
var ErrInputFileNotFound = fmt.Errorf("input file not found")

// ErrOutputFileNotFound predefined error.
var ErrOutputFileNotFound = fmt.Errorf("output file not found")

// Task is interface.
type Task interface {
	Run(inputData string, id int) string
}

// TaskFn type - string -> string function
// TaskFn satisfies Task interface.
type TaskFn func(inputData string) string

// Run task (just simple call fn, without passing id)
// if id of task is essential use Task interface itself.
func (fn TaskFn) Run(inputData string, id int) string {
	return fn(inputData)
}

// TaskTestResult is result of running concrete test on current task.
type TaskTestResult struct {
	ID             int    // ID of test
	Ok             bool   // Fail or Ok test was
	Run            bool   // Has test run yet
	Skipped        bool   // Has test be skipped intentionally
	Expected       string // Expected result
	Actual         string // Actual result
	Input          string // Input of task
	InputFilePath  string // File path of file where is test input
	OutputFilePath string // File path of file where is test output (expected result of task)
	Err            error  // Error of reading input or output file
}

// TaskTester runs bunch of tests from directory on current task.
type TaskTester struct {
	task Task
}

// NewTaskTesterFn constructs of task tester for current function - task (TaskFn).
func NewTaskTesterFn(task TaskFn) *TaskTester {
	return &TaskTester{
		task: task,
	}
}

// NewTaskTesterFn constructs of task tester for current task (Task interface).
func NewTaskTester(task Task) *TaskTester {
	return &TaskTester{
		task: task,
	}
}

// RunDir run all tests in directory for current task, return list of results.
func (t *TaskTester) RunDir(dir string) []*TaskTestResult {
	return t.RunDirWithSkipped(dir, nil)
}

// RunDirWithSkipped run not all tests in directory for current task, return list of results
// Which test to skip determined by skip predicate.
func (t *TaskTester) RunDirWithSkipped(dir string, skip func(inputData string) bool) []*TaskTestResult {
	results, err := t.scanDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// results map is filled up - run tests
	for _, result := range results {
		if result.InputFilePath == "" {
			result.Err = ErrInputFileNotFound
			continue
		}

		if result.OutputFilePath == "" {
			result.Err = ErrOutputFileNotFound
			continue
		}

		content, err := ioutil.ReadFile(result.InputFilePath)
		if err != nil {
			result.Err = err
			continue
		}

		input := strings.TrimSpace(string(content))

		if skip != nil && skip(input) {
			result.Skipped = true
			continue
		}

		// Save input
		result.Input = input

		content, err = ioutil.ReadFile(result.OutputFilePath)
		if err != nil {
			result.Err = err
			continue
		}

		expected := strings.TrimSpace(string(content))
		expected = normalizaLineEnding(expected)

		// Run task itself
		output := strings.TrimSpace(t.task.Run(input, result.ID))
		output = normalizaLineEnding(output)

		// Mark that task has been run
		result.Run = true

		// Save expected and actual
		result.Expected = expected
		result.Actual = output

		// Fail or Ok of run task
		result.Ok = expected == output
	}

	return resultsMapToSlice(results)
}

// scanDir is protected helper that scans directory for special test files and fill results map.
func (t *TaskTester) scanDir(dir string) (map[int]*TaskTestResult, error) {
	taskResults := make(map[int]*TaskTestResult)

	stat, statErr := os.Stat(dir)
	if os.IsNotExist(statErr) {
		return nil, fmt.Errorf("%s not exists", dir)
	}

	if !stat.IsDir() {
		return nil, fmt.Errorf("%s not directory", dir)
	}

	walkErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		name := info.Name()

		parts := strings.Split(name, ".")
		if len(parts) < 3 || parts[0] != "test" {
			return nil
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil
		}

		result, ok := taskResults[id]
		if !ok {
			result = &TaskTestResult{
				ID: id,
			}
		}

		switch parts[2] {
		case "in":
			result.InputFilePath = path
		case "out":
			result.OutputFilePath = path
		default:
			return nil
		}

		taskResults[id] = result

		return nil
	})

	return taskResults, walkErr
}

func resultsMapToSlice(results map[int]*TaskTestResult) []*TaskTestResult {
	resultList := make([]*TaskTestResult, len(results))

	i := 0

	for _, result := range results {
		resultList[i] = result
		i++
	}

	sort.Slice(resultList, func(i, j int) bool {
		return resultList[i].ID < resultList[j].ID
	})

	return resultList
}

func normalizaLineEnding(str string) string {
	str = strings.Replace(str, "\r\n", "\n", -1)
	str = strings.Replace(str, "\r", "\n", -1)

	return str
}
