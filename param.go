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
type AssertWrapper[O OutputT] func(t *testing.T, expected, actual O)

// GroupTest iterates over each individual TestCase, calling the UnitWrapper on the input, and then
// passing the TestCase's Output (as expected) and the wrapper's Output (as actual) to the AssertWrapper
func GroupTest[I InputT, O OutputT](t *testing.T, testCases []TestCase[I, O], w UnitWrapper[I, O], a AssertWrapper[O]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			a(t, tc.Expected, w(tc.Input))
		})
	}
}

// NaiveGenericAssert is a simple default assertion wrapper that directly compares expected and actual
// output values with a simple equality check
func GenericEqualityAssert[O comparable](t *testing.T, expected, actual O) {
	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func GenericSliceEqualityAssert[O comparable](t *testing.T, expected, actual []O) {
	if len(expected) != len(actual) {
		t.Fatalf("expected %d elements in slice, got %d", len(expected), len(actual))
	}
	for i := range expected {
		GenericEqualityAssert(t, expected[i], actual[i])
	}
}

// DefaultGroupTest uses NaiveGenericAssert and relies upon the Output Type satisfying the standard
// `comparable` type constraint
func DefaultGroupTest[I InputT, O comparable](t *testing.T, testCases []TestCase[I, O], w UnitWrapper[I, O]) {
	GroupTest(t, testCases, w, GenericEqualityAssert[O])
}

func SliceGroupTest[I InputT, O comparable](t *testing.T, testCases []TestCase[I, []O], w UnitWrapper[I, []O]) {
	GroupTest(t, testCases, w, GenericSliceEqualityAssert[O])
}
