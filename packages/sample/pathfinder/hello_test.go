package main

import (
	"reflect"
	"testing"
)

func TestFindPath(t *testing.T) {
	g := [][]string{
		{"0", "1", "0", "0"},
		{"0", "1", "0", "1"},
		{"0", "0", "0", "1"},
	}

	start := Position{0, 0}
	end := Position{0, 3}

	actual := FindPath(start, end, g)
	expected := []string{"(0, 0)", "(1, 0)", "(2, 0)", "(2, 1)", "(2, 2)", "(1, 2)", "(0, 2)", "(0, 3)"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Path does not match expected result\nExpected:\n%v\nReceived:\n%v", expected, actual)
	}
}
