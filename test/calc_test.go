package test

import (
	"calculator/internal/services/calc"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc_BaseFunctions(t *testing.T) {
	testTable := []struct {
		in       string
		expected string
	}{
		{
			in:       "2412 + 3526",
			expected: "5938",
		},
		{
			in:       "7251 - 51934",
			expected: "-44683",
		},
		{
			in:       "7812 * 62",
			expected: "484344",
		},
		{
			in:       "36713 / 13",
			expected: "2824.076923076923",
		},
		{
			in:       "36713 ^ 4",
			expected: "1.8166844430450081e+18",
		},
		{
			in:       "rt 4 1892",
			expected: "6.595235124077313",
		},
		{
			in:       "log 12 97312",
			expected: "4.622176688793115",
		},
	}

	for _, tt := range testTable {
		ans, err := calc.CalculateExpr(tt.in)

		t.Logf("Calling CalculateExpr(%v), result %v\n", tt.in, ans)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, tt.expected, ans,
			fmt.Sprintf("expected %s, got %s", tt.expected, ans))
	}
}

func TestCalc_ExtraExpressions(t *testing.T) {
	testTable := []struct {
		in       string
		expected string
	}{
		{
			in:       "( 12673 + 32 ) * 67",
			expected: "851235",
		},
		{
			in:       "( 2 * 3 * 2 / 10 ) + 83 * 89",
			expected: "7388.2",
		},
		{
			in:       "( ( 123 + 7632 ) ^ 2 - ( 276 + 172 ) ^ 3 ) ^ 2",
			expected: "8.86572479984689e+14",
		},
		{
			in:       "( ( 124 * 6712 - 6732 / 76 ) / ( 81 ^ ( 2 / 3 ) * 8921 ) ) ^ ( 21 / 24 )",
			expected: "4.076657338152463",
		},
		{
			in:       "2 ^ ( rt ( 2 ^ 4 ) log ( 2 ^ 5 ) ( 12763 * 63 ) )",
			expected: "2.127546300089071",
		},
		{
			in:       "( ( 1 + 2 ) ^ 2 - rt ( 3 ^ 2 ) ( log ( 2 ^ 4 ) ( 4 * 8 ) ) ) * 8",
			expected: "63.79917081234479",
		},
	}

	for _, tt := range testTable {
		ans, err := calc.CalculateExpr(tt.in)

		t.Logf("Calling CalculateExpr(%v), result %v\n", tt.in, ans)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, tt.expected, ans,
			fmt.Sprintf("expected %s, got %s", tt.expected, ans))
	}
}
