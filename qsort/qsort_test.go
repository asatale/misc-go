package qsort

import (
	"testing"
)

type testCase struct {
	input  []interface{}
	output []interface{}
}

func equal(a, b []interface{}, cmp Compare) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if cmp(v, b[i]) != 0 {
			return false
		}
	}
	return true
}

func runTestcases(tests []testCase, cmp Compare, t *testing.T) {
	for _, test := range tests {
		Sort(test.input, cmp)
		if !equal(test.input, test.output, cmp) {
			t.Errorf("Failed. Expected: %v, Got: %v", test.output, test.input)
		}
	}
}

func TestIntegerSort(t *testing.T) {

	testCases := []testCase{
		{
			input:  []interface{}{4, 3, 2, 1},
			output: []interface{}{1, 2, 3, 4},
		},
		{
			input:  []interface{}{1, 2, 3, 4},
			output: []interface{}{1, 2, 3, 4},
		},
		{
			input:  []interface{}{2, 1},
			output: []interface{}{1, 2},
		},
		{
			input:  []interface{}{3, 2, 4, 1, 5, 5, 6, 1},
			output: []interface{}{1, 1, 2, 3, 4, 5, 5, 6},
		},
		{
			input:  []interface{}{1, 1, 1, 1, 1, 1, 1, 1},
			output: []interface{}{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			input:  []interface{}{1},
			output: []interface{}{1},
		},
		{
			input:  []interface{}{},
			output: []interface{}{},
		},
	}
	runTestcases(testCases,
		func(a, b interface{}) int {
			v1 := a.(int)
			v2 := b.(int)
			if v1 == v2 {
				return 0
			} else if v1 < v2 {
				return -1
			}
			return 1
		},
		t)
}

func TestStringSort(t *testing.T) {
	testCases := []testCase{
		{
			input:  []interface{}{"b", "a"},
			output: []interface{}{"a", "b"},
		},
		{
			input:  []interface{}{"d", "c", "b", "a"},
			output: []interface{}{"a", "b", "c", "d"},
		},
		{
			input:  []interface{}{"A", "a", "AA", "aa"},
			output: []interface{}{"A", "AA", "a", "aa"},
		},
		{
			input:  []interface{}{"da", "ac", "zb", "a"},
			output: []interface{}{"a", "ac", "da", "zb"},
		},
	}
	runTestcases(testCases,
		func(a, b interface{}) int {
			v1 := a.(string)
			v2 := b.(string)
			if v1 == v2 {
				return 0
			} else if v1 < v2 {
				return -1
			}
			return 1
		},
		t)
}
