// +build !actual

package fibonacci

import (
	"math/big"
	"testing"
)

func TestPow0(t *testing.T) {
	x := big.NewFloat(10)

	Pow(x, 0)

	expected := big.NewFloat(1)

	if expected.String() != x.String() {
		t.Errorf("unexpected %s instread of %s", x.String(), expected.String())
	}
}

func TestPow1(t *testing.T) {
	x := big.NewFloat(10)

	Pow(x, 1)

	expected := big.NewFloat(10)

	if expected.String() != x.String() {
		t.Errorf("unexpected %s instread of %s", x.String(), expected.String())
	}
}

func TestPow2(t *testing.T) {
	x := big.NewFloat(10)

	Pow(x, 2)

	expected := big.NewFloat(100)

	if expected.String() != x.String() {
		t.Errorf("unexpected %s instread of %s", x.String(), expected.String())
	}
}
