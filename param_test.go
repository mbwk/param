package param_test

import (
	"strings"
	"testing"

	"github.com/mbwk/param"
)

func TestTrimSpace(t *testing.T) {
	testCases := []param.TestCase[string, string]{
		{Name: "basic", Input: "one ", Expected: "one"},
		{Name: "noop", Input: "one", Expected: "one"},
		{Name: "empty", Input: "", Expected: ""},
		{Name: "sentence", Input: "the quick brown fox", Expected: "the quick brown fox"},
		{Name: "sentence margin", Input: "   the quick brown fox     ", Expected: "the quick brown fox"},
		{Name: "full width spaces", Input: "　日本語　", Expected: "日本語"},
	}

	param.DefaultGroupTest(t, testCases, strings.TrimSpace)
}

func TestTrim(t *testing.T) {
	testCases := []param.TestCase[string, string]{
		{Name: "basic", Input: "one ", Expected: "one"},
		{Name: "noop", Input: "one", Expected: "one"},
		{Name: "empty", Input: "", Expected: ""},
		{Name: "sentence", Input: "the quick brown fox", Expected: "the quick brown fox"},
		{Name: "sentence margin", Input: "   the quick brown fox     ", Expected: "the quick brown fox"},
		{Name: "full width spaces", Input: "　日本語　", Expected: "日本語"},
	}

	param.DefaultGroupTest(t, testCases, func(input string) string {
		return strings.Trim(input, " 　")
	})
}

func TestSplit(t *testing.T) {
	type splitInput struct {
		S string
		Sep string
	}
	testCases := []param.TestCase[splitInput, []string]{
		{
			Name: "simple",
			Input: splitInput{
				S: "one two three",
				Sep: " ",
			},
			Expected: []string{"one", "two", "three"},
		},
		{
			Name: "single",
			Input: splitInput{
				S: "one",
				Sep: " ",
			},
			Expected: []string{"one"},
		},
		{
			Name: "empty separator",
			Input: splitInput{
				S: "one",
				Sep: "",
			},
			Expected: []string{"o","n","e"},
		},
	}

	param.GroupTest(t, testCases, func(i splitInput) []string {
		return strings.Split(i.S, i.Sep)
	}, func (t *testing.T, expected []string, actual []string) {
		if len(expected) != len(actual) {
			t.Errorf("slice length mismatch %d != %d", len(expected), len(actual))
		}
		for i := range expected {
			if expected[i] != actual[i] {
				t.Errorf("slice contents mismatch `%s` != `%s`", expected[i], actual[i])
			}
		}
	})
}
