package cmp_test

import (
	"testing"

	"math/big"

	"github.com/reiver/go-cmp"
)

func TestGreaterThan_bigInt(t *testing.T) {

	tests := []struct{
		Val1 cmp.Comparer[*big.Int]
		Val2 *big.Int
		Expected bool
	}{
		{
			Val1: big.NewInt(-3),
			Val2: big.NewInt(-3),
			Expected: false,
		},
		{
			Val1: big.NewInt(-3),
			Val2: big.NewInt(-2),
			Expected: false,
		},
		{
			Val1: big.NewInt(-3),
			Val2: big.NewInt(-1),
			Expected: false,
		},
		{
			Val1: big.NewInt(-3),
			Val2: big.NewInt(0),
			Expected: false,
		},
		{
			Val1: big.NewInt(-3),
			Val2: big.NewInt(1),
			Expected: false,
		},
		{
			Val1: big.NewInt(-3),
			Val2: big.NewInt(2),
			Expected: false,
		},
		{
			Val1: big.NewInt(-3),
			Val2: big.NewInt(3),
			Expected: false,
		},









		{
			Val1: big.NewInt(-2),
			Val2: big.NewInt(-3),
			Expected: true,
		},
		{
			Val1: big.NewInt(-2),
			Val2: big.NewInt(-2),
			Expected: false,
		},
		{
			Val1: big.NewInt(-2),
			Val2: big.NewInt(-1),
			Expected: false,
		},
		{
			Val1: big.NewInt(-2),
			Val2: big.NewInt(0),
			Expected: false,
		},
		{
			Val1: big.NewInt(-2),
			Val2: big.NewInt(1),
			Expected: false,
		},
		{
			Val1: big.NewInt(-2),
			Val2: big.NewInt(2),
			Expected: false,
		},
		{
			Val1: big.NewInt(-2),
			Val2: big.NewInt(3),
			Expected: false,
		},









		{
			Val1: big.NewInt(-1),
			Val2: big.NewInt(-3),
			Expected: true,
		},
		{
			Val1: big.NewInt(-1),
			Val2: big.NewInt(-2),
			Expected: true,
		},
		{
			Val1: big.NewInt(-1),
			Val2: big.NewInt(-1),
			Expected: false,
		},
		{
			Val1: big.NewInt(-1),
			Val2: big.NewInt(0),
			Expected: false,
		},
		{
			Val1: big.NewInt(-1),
			Val2: big.NewInt(1),
			Expected: false,
		},
		{
			Val1: big.NewInt(-1),
			Val2: big.NewInt(2),
			Expected: false,
		},
		{
			Val1: big.NewInt(-1),
			Val2: big.NewInt(3),
			Expected: false,
		},









		{
			Val1: big.NewInt(0),
			Val2: big.NewInt(-3),
			Expected: true,
		},
		{
			Val1: big.NewInt(0),
			Val2: big.NewInt(-2),
			Expected: true,
		},
		{
			Val1: big.NewInt(0),
			Val2: big.NewInt(-1),
			Expected: true,
		},
		{
			Val1: big.NewInt(0),
			Val2: big.NewInt(0),
			Expected: false,
		},
		{
			Val1: big.NewInt(0),
			Val2: big.NewInt(1),
			Expected: false,
		},
		{
			Val1: big.NewInt(0),
			Val2: big.NewInt(2),
			Expected: false,
		},
		{
			Val1: big.NewInt(0),
			Val2: big.NewInt(3),
			Expected: false,
		},









		{
			Val1: big.NewInt(1),
			Val2: big.NewInt(-3),
			Expected: true,
		},
		{
			Val1: big.NewInt(1),
			Val2: big.NewInt(-2),
			Expected: true,
		},
		{
			Val1: big.NewInt(1),
			Val2: big.NewInt(-1),
			Expected: true,
		},
		{
			Val1: big.NewInt(1),
			Val2: big.NewInt(0),
			Expected: true,
		},
		{
			Val1: big.NewInt(1),
			Val2: big.NewInt(1),
			Expected: false,
		},
		{
			Val1: big.NewInt(1),
			Val2: big.NewInt(2),
			Expected: false,
		},
		{
			Val1: big.NewInt(1),
			Val2: big.NewInt(3),
			Expected: false,
		},









		{
			Val1: big.NewInt(2),
			Val2: big.NewInt(-3),
			Expected: true,
		},
		{
			Val1: big.NewInt(2),
			Val2: big.NewInt(-2),
			Expected: true,
		},
		{
			Val1: big.NewInt(2),
			Val2: big.NewInt(-1),
			Expected: true,
		},
		{
			Val1: big.NewInt(2),
			Val2: big.NewInt(0),
			Expected: true,
		},
		{
			Val1: big.NewInt(2),
			Val2: big.NewInt(1),
			Expected: true,
		},
		{
			Val1: big.NewInt(2),
			Val2: big.NewInt(2),
			Expected: false,
		},
		{
			Val1: big.NewInt(2),
			Val2: big.NewInt(3),
			Expected: false,
		},









		{
			Val1: big.NewInt(3),
			Val2: big.NewInt(-3),
			Expected: true,
		},
		{
			Val1: big.NewInt(3),
			Val2: big.NewInt(-2),
			Expected: true,
		},
		{
			Val1: big.NewInt(3),
			Val2: big.NewInt(-1),
			Expected: true,
		},
		{
			Val1: big.NewInt(3),
			Val2: big.NewInt(0),
			Expected: true,
		},
		{
			Val1: big.NewInt(3),
			Val2: big.NewInt(1),
			Expected: true,
		},
		{
			Val1: big.NewInt(3),
			Val2: big.NewInt(2),
			Expected: true,
		},
		{
			Val1: big.NewInt(3),
			Val2: big.NewInt(3),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		var actual   bool = cmp.GreaterThan(test.Val1, test.Val2)
		var expected bool = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual greater-than-value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VAL-1: %v", test.Val1)
			t.Logf("VAL-2: %v", test.Val2)
			continue
		}
	}
}
