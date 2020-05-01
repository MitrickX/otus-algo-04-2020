package tester

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Predefined errors
var TestErrorInputFileNotFound = fmt.Errorf("input file not found")
var TestErrorOutputFileNotFound = fmt.Errorf("output file not found")

// Task type - string -> string function
type Task func(inputData string) string

// TaskTestResult is result of running concrete test on current task
type TaskTestResult struct {
	Id             int    // Id of test
	Ok             bool   // Fail or Ok test was
	Run            bool   // Has test run yet
	InputFilePath  string // File path of file where is test input
	OutputFilePath string // File path of file where is test output (expected result of task)
	Err            error  // Error of reading input or output file
}

// TaskTester runs bunch of tests from directory on current task
type TaskTester struct {
	task Task
}

// NewTaskTester constructs of task tester for current task
func NewTaskTester(task Task) *TaskTester {
	return &TaskTester{
		task: task,
	}
}

// RunDir run all tests in directory for current task, return list of results
func (t *TaskTester) RunDir(dir string) []*TaskTestResult {

	results, err := t.scanDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// results map is filled up - run tests
	for _, result := range results {

		if result.InputFilePath == "" {
			result.Err = TestErrorInputFileNotFound
			continue
		}

		if result.OutputFilePath == "" {
			result.Err = TestErrorOutputFileNotFound
			continue
		}

		content, err := ioutil.ReadFile(result.InputFilePath)
		if err != nil {
			result.Err = err
			continue
		}

		input := strings.TrimSpace(string(content))

		content, err = ioutil.ReadFile(result.OutputFilePath)
		if err != nil {
			result.Err = err
			continue
		}

		expected := strings.TrimSpace(string(content))

		// Run task itself
		output := strings.TrimSpace(t.task(input))

		// Mark that task has been run
		result.Run = true

		// Fail or Ok of run task
		result.Ok = expected == output
	}

	resultList := make([]*TaskTestResult, len(results))

	i := 0
	for _, result := range results {
		resultList[i] = result
		i++
	}

	return resultList
}

// scanDir is protected helper that scans directory for special test files and fill results map
func (t *TaskTester) scanDir(dir string) (map[int]*TaskTestResult, error) {
	taskResults := make(map[int]*TaskTestResult)

	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("%s not exists", dir)
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("%s not directory", dir)
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
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
				Id: id,
			}
		}

		if parts[2] == "in" {
			result.InputFilePath = path
		} else if parts[2] == "out" {
			result.OutputFilePath = path
		} else {
			return nil
		}

		taskResults[id] = result

		return nil
	})

	return taskResults, nil
}
