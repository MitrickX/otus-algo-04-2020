// +build !actual

package spells

import (
	"bytes"
	"fmt"
	"math"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

type spellTask struct {
	spells map[string]Spell // spell string expression to spell function
}

func newSpellTask(spells map[string]Spell) spellTask {
	return spellTask{
		spells: spells,
	}
}

func (s spellTask) Run(inputData string, _ int) string {
	expr := strings.TrimSpace(inputData)
	spell, ok := s.spells[expr]

	if !ok {
		return fmt.Sprintf("spell '%s' not found", expr)
	}

	writer := bytes.NewBufferString("")
	Draw(spell, writer)

	return strings.TrimSpace(writer.String())
}

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

func TestDraw(t *testing.T) {
	// map of spells
	spells := map[string]Spell{
		"x < y": func(x, y int) bool {
			return x < y
		},
		"x == y": func(x, y int) bool {
			return x == y
		},
		"24-x == y": func(x, y int) bool {
			return 24-x == y
		},
		"30-x > y": func(x, y int) bool {
			return 30-x > y
		},
		"x == floor(y/2)": func(x, y int) bool {
			return x == int(math.Floor(float64(y)/2))
		},
		"x < 10 || y < 10": func(x, y int) bool {
			return x < 10 || y < 10
		},
		"x > 15 && y > 15": func(x, y int) bool {
			return x > 15 && y > 15
		},
		"x*y == 0": func(x, y int) bool {
			return x*y == 0 // same as x == 0 || y == 0
		},
		"abs(x-y) > 10": func(x, y int) bool {
			return math.Abs(float64(x-y)) > 10 // same as x - y > 10 || y - x > 10
		},
		"y >= x+1 && y <= 2*x+1": func(x, y int) bool {
			return y >= x+1 && y <= 2*x+1
		},
		"x == 1 || (24-x) == 1 || y == 1 || (24-y) == 1": func(x, y int) bool {
			return x == 1 || (24-x) == 1 || y == 1 || (24-y) == 1
		},
		"x*x+y*y <= 400": func(x, y int) bool {
			return x*x+y*y <= 400
		},
		"abs(24-x-y) <= 4": func(x, y int) bool {
			return math.Abs(float64(24-x-y)) <= 4
		},
	}

	spellTask := newSpellTask(spells)
	taskTester := tester.NewTaskTester(spellTask)

	dataDir := getDataDir()
	resultList := taskTester.RunDir(dataDir)

	for _, result := range resultList {
		if !result.Run {
			t.Errorf("Test #%d has not been run because of %s", result.ID, result.Err)
		} else if !result.Ok {
			t.Errorf("Test #%d (%s) is fail\nUnexpected:\n%s\n\nInstread of:\n%s",
				result.ID,
				result.Input,
				result.Actual,
				result.Expected)
		}
	}
}
