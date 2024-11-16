package math_test // black box testing

import (
	"os"
	"test-example/math"
	"testing"
)

/*
// Single test

	func TestMinInt(t *testing.T) {
		result := math.MinInt(10, 5)
		if result != 5 {
			t.Errorf("expected: %d, got: %d", 5, result)
		}
	}
*/

func TestMain(m *testing.M) {
	exitCode := m.Run()

	os.Exit(exitCode)
}

// Table test - have multiple test cases

func TestMinInt(t *testing.T) {
	t.Cleanup(func() {
		// cleanup if neededS
	})
	testCases := []struct { // anonyms struct
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "100 is less than 1000",
			a:        100,
			b:        1000,
			expected: 100,
		},
		{
			name:     "-2 is less than -1",
			a:        -1,
			b:        -2,
			expected: -2,
		},
		{
			name:     "69 is less than 6969",
			a:        69,
			b:        6969,
			expected: 69,
		},
		{
			name:     "-100 is less than 100",
			a:        -100,
			b:        100,
			expected: -100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := math.MinInt(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected: %d, got: %d", tc.expected, result)
			}
		})
	}

}

func TestMaxInt(t *testing.T) {
	result := math.MaxInt(10, 5)
	if result != 10 {
		t.Errorf("expected: %d, got: %d", 10, result)
	}
}
