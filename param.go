package param

import "testing"

type InputT any
type OutputT any

// TestCase encapsulates structure test cases which act as inputs to
// parameterized tests
type TestCase[I InputT, O OutputT] struct {
	Name     string
	Input    I
	Expected O
}

// UnitWrapper converts the inputs and outputs from a given tested unit
// to and from a form understood by TestCase
type UnitWrapper[I InputT, O OutputT] func(i I) O

// AssertWrapper encapsulates the assertion portion of a test, and handles comparisons between
// expected and actual test outputs
type AssertWrapper[O OutputT] func(t *testing.T, expected O, actual O)

// GroupTest 
func GroupTest[I InputT, O OutputT](t *testing.T, testCases []TestCase[I, O], w UnitWrapper[I, O], a AssertWrapper[O]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			a(t, tc.Expected, w(tc.Input))
		})
	}
}

// NaiveGenericAssert is a simple default assertion wrapper that directly compares expected and actual
// output values with a simple equality check
func NaiveGenericAssert[O comparable](t *testing.T, expected O, actual O) {
	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

// DefaultGroupTest uses NaiveGenericAssert and relies upon the Output Type satisfying the standard
// `comparable` type constraint
func DefaultGroupTest[I InputT, O comparable](t *testing.T, testCases []TestCase[I, O], w UnitWrapper[I, O]) {
	GroupTest(t, testCases, w, NaiveGenericAssert[O])
}
