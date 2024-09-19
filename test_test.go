package cmp_test

import (
	"testing"

	"math/big"

	"github.com/reiver/go-cmp"
)

func TestTest_LessThan(t *testing.T) {

	var x *big.Int = big.NewInt(-5)
	var y *big.Int = big.NewInt(5)

	{
		var actual   bool = cmp.LessThan(x, y)
		var expected bool = true

		if expected != actual {
			t.Errorf("The actual less-than-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X < Y -> %v < %v", x, y)
		}
	}

	{
		var actual   bool = cmp.Equal(x, y)
		var expected bool = false

		if expected != actual {
			t.Errorf("The actual equal-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X == Y -> %v == %v", x, y)
		}
	}

	{
		var actual   bool = cmp.GreaterThan(x, y)
		var expected bool = false

		if expected != actual {
			t.Errorf("The actual greater-than-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X > Y -> %v > %v", x, y)
		}
	}
}

func TestTest_Equal(t *testing.T) {

	var x *big.Int = big.NewInt(5)
	var y *big.Int = big.NewInt(5)

	{
		var actual   bool = cmp.LessThan(x, y)
		var expected bool = false

		if expected != actual {
			t.Errorf("The actual less-than-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X < Y -> %v < %v", x, y)
		}
	}

	{
		var actual   bool = cmp.Equal(x, y)
		var expected bool = true

		if expected != actual {
			t.Errorf("The actual equal-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X == Y -> %v == %v", x, y)
		}
	}

	{
		var actual   bool = cmp.GreaterThan(x, y)
		var expected bool = false

		if expected != actual {
			t.Errorf("The actual greater-than-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X > Y -> %v > %v", x, y)
		}
	}
}

func TestTest_GreaterThan(t *testing.T) {

	var x *big.Int = big.NewInt(5)
	var y *big.Int = big.NewInt(-5)

	{
		var actual   bool = cmp.LessThan(x, y)
		var expected bool = false

		if expected != actual {
			t.Errorf("The actual less-than-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X < Y -> %v < %v", x, y)
		}
	}

	{
		var actual   bool = cmp.Equal(x, y)
		var expected bool = false

		if expected != actual {
			t.Errorf("The actual equal-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X == Y -> %v == %v", x, y)
		}
	}

	{
		var actual   bool = cmp.GreaterThan(x, y)
		var expected bool = true

		if expected != actual {
			t.Errorf("The actual greater-than-value is not what was expected.")
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("X > Y -> %v > %v", x, y)
		}
	}
}
