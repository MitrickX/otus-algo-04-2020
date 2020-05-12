package gcd

import (
	"math/big"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/MitrickX/otus-algo-04-2020/tester"
)

func TestSubtract(t *testing.T) {
	testGCD(t, Subtract, "Subtract", func(inputData string) bool {
		a, b := parseInputData(inputData)
		limit := big.NewInt(8564657687654654657)
		return a.Cmp(limit) > 0 || b.Cmp(limit) > 0
	})
}

func TestRest(t *testing.T) {
	testGCD(t, Rest, "Rest", nil)
}

func TestBinary(t *testing.T) {
	testGCD(t, Binary, "Binary", nil)
}

// BENCHMARKS

// x=1, y=100.
func BenchmarkSubtract_1_100(b *testing.B) {
	x := big.NewInt(1)
	y := big.NewInt(100)

	for i := 0; i < b.N; i++ {
		_ = Subtract(x, y)
	}
}

// x=1, y=100.
func BenchmarkRest_1_100(b *testing.B) {
	x := big.NewInt(1)
	y := big.NewInt(100)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=1, y=100.
func BenchmarkBinary_1_100(b *testing.B) {
	x := big.NewInt(1)
	y := big.NewInt(100)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=20, y=30.
func BenchmarkSubtract_20_30(b *testing.B) {
	x := big.NewInt(20)
	y := big.NewInt(30)

	for i := 0; i < b.N; i++ {
		_ = Subtract(x, y)
	}
}

// x=20, y=30.
func BenchmarkRest_20_30(b *testing.B) {
	x := big.NewInt(20)
	y := big.NewInt(30)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=20, y=30.
func BenchmarkBinary_20_30(b *testing.B) {
	x := big.NewInt(20)
	y := big.NewInt(30)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=2, y=123456789.
func BenchmarkSubtract_2_123456789(b *testing.B) {
	x := big.NewInt(2)
	y := big.NewInt(123456789)

	for i := 0; i < b.N; i++ {
		_ = Subtract(x, y)
	}
}

// x=2, y=123456789.
func BenchmarkRest_2_123456789(b *testing.B) {
	x := big.NewInt(2)
	y := big.NewInt(123456789)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=2, y=123456789.
func BenchmarkBinary_2_123456789(b *testing.B) {
	x := big.NewInt(2)
	y := big.NewInt(123456789)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=12, y=1234567890.
func BenchmarkSubtract_12_1234567890(b *testing.B) {
	x := big.NewInt(12)
	y := big.NewInt(1234567890)

	for i := 0; i < b.N; i++ {
		_ = Subtract(x, y)
	}
}

// x=12, y=1234567890.
func BenchmarkRest_12_1234567890(b *testing.B) {
	x := big.NewInt(12)
	y := big.NewInt(1234567890)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=12, y=1234567890.
func BenchmarkBinary_12_1234567890(b *testing.B) {
	x := big.NewInt(12)
	y := big.NewInt(1234567890)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=123456789, y=9876543210.
func BenchmarkSubtract_123456789_9876543210(b *testing.B) {
	x := big.NewInt(123456789)
	y := big.NewInt(9876543210)

	for i := 0; i < b.N; i++ {
		_ = Subtract(x, y)
	}
}

// x=123456789, y=9876543210.
func BenchmarkRest_123456789_9876543210(b *testing.B) {
	x := big.NewInt(123456789)
	y := big.NewInt(9876543210)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=123456789, y=9876543210.
func BenchmarkBinary_123456789_9876543210(b *testing.B) {
	x := big.NewInt(123456789)
	y := big.NewInt(9876543210)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=997, y=1073676287.
func BenchmarkSubtract_997_1073676287(b *testing.B) {
	x := big.NewInt(997)
	y := big.NewInt(1073676287)

	for i := 0; i < b.N; i++ {
		_ = Subtract(x, y)
	}
}

// x=997, y=1073676287.
func BenchmarkRest_997_1073676287(b *testing.B) {
	x := big.NewInt(997)
	y := big.NewInt(1073676287)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=997, y=1073676287.
func BenchmarkBinary_997_1073676287(b *testing.B) {
	x := big.NewInt(997)
	y := big.NewInt(1073676287)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=8564657687654654657, y=1298074214633706835075030044377087.
func BenchmarkRest_8564657687654654657_1298074214633706835075030044377087(b *testing.B) {
	x, _ := new(big.Int).SetString("8564657687654654657", 10)
	y, _ := new(big.Int).SetString("1298074214633706835075030044377087", 10)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=8564657687654654657, y=1298074214633706835075030044377087.
func BenchmarkBinary_8564657687654654657_1298074214633706835075030044377087(b *testing.B) {
	x, _ := new(big.Int).SetString("8564657687654654657", 10)
	y, _ := new(big.Int).SetString("1298074214633706835075030044377087", 10)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=162259276829213363391578010288127, y=123426017006182806728593424683999798008235734137469123231828679.
func BenchmarkRest_162259276829213363391578010288127_123426017006182806728593424683999798008235734137469123231828679(b *testing.B) { //nolint:lll
	x, _ := new(big.Int).SetString("162259276829213363391578010288127", 10)
	y, _ := new(big.Int).SetString("123426017006182806728593424683999798008235734137469123231828679", 10)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=162259276829213363391578010288127, y=123426017006182806728593424683999798008235734137469123231828679.
func BenchmarkBinary_162259276829213363391578010288127_123426017006182806728593424683999798008235734137469123231828679(b *testing.B) { //nolint:lll
	x, _ := new(big.Int).SetString("162259276829213363391578010288127", 10)
	y, _ := new(big.Int).SetString("123426017006182806728593424683999798008235734137469123231828679", 10)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=162259276829213363391578010288126, y=123426017006182806728593424683999798008235734137469123231828678.
func BenchmarkRest_162259276829213363391578010288126_123426017006182806728593424683999798008235734137469123231828678(b *testing.B) { //nolint:lll
	x, _ := new(big.Int).SetString("162259276829213363391578010288126", 10)
	y, _ := new(big.Int).SetString("123426017006182806728593424683999798008235734137469123231828678", 10)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=162259276829213363391578010288126, y=123426017006182806728593424683999798008235734137469123231828678.
func BenchmarkBinary_162259276829213363391578010288126_123426017006182806728593424683999798008235734137469123231828678(b *testing.B) { //nolint:lll
	x, _ := new(big.Int).SetString("162259276829213363391578010288126", 10)
	y, _ := new(big.Int).SetString("123426017006182806728593424683999798008235734137469123231828678", 10)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
	}
}

// x=608281864034267560872252163321295376887552831379210240000000000,
// y=30414093201713378043612608166064768844377641568960512000000000000.
func BenchmarkRest_608281864034267560872252163321295376887552831379210240000000000_30414093201713378043612608166064768844377641568960512000000000000(b *testing.B) { //nolint:lll
	x, _ := new(big.Int).SetString("608281864034267560872252163321295376887552831379210240000000000", 10)
	y, _ := new(big.Int).SetString("30414093201713378043612608166064768844377641568960512000000000000", 10)

	for i := 0; i < b.N; i++ {
		_ = Rest(x, y)
	}
}

// x=608281864034267560872252163321295376887552831379210240000000000,
// y=30414093201713378043612608166064768844377641568960512000000000000.
func BenchmarkBinary_608281864034267560872252163321295376887552831379210240000000000_30414093201713378043612608166064768844377641568960512000000000000(b *testing.B) { //nolint:lll
	x, _ := new(big.Int).SetString("608281864034267560872252163321295376887552831379210240000000000", 10)
	y, _ := new(big.Int).SetString("30414093201713378043612608166064768844377641568960512000000000000", 10)

	for i := 0; i < b.N; i++ {
		_ = Binary(x, y)
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

func parseInputData(inputData string) (*big.Int, *big.Int) {
	inputData = strings.Replace(inputData, "\r\n", "\n", -1)
	numbers := strings.Split(inputData, "\n")

	a := new(big.Int)
	b := new(big.Int)

	a.SetString(numbers[0], 10)
	b.SetString(numbers[1], 10)

	return a, b
}

func convertToTask(fn func(*big.Int, *big.Int) *big.Int) func(string) string {
	return func(inputData string) string {
		a, b := parseInputData(inputData)
		result := fn(a, b)

		return result.String()
	}
}

func testGCD(t *testing.T, fn func(*big.Int, *big.Int) *big.Int, name string, skip func(string) bool) {
	task := convertToTask(fn)
	taskTester := tester.NewTaskTesterFn(task)
	resultList := taskTester.RunDirWithSkipped(getInputsDir(), skip)

	for _, result := range resultList {
		// test intentionally skipped
		if result.Skipped {
			continue
		}

		// test not run for some reason
		if !result.Run {
			t.Errorf("%s - test #%d has not been run because of %s", name, result.ID, result.Err)
			continue
		}

		if !result.Ok {
			t.Errorf("%s - test #%d is fail, unexpected `%s` instread of `%s`", name, result.ID, result.Actual, result.Expected)
			break
		}
	}
}
