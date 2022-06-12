package main

import (
	"reflect"
	"testing"
)

func TestFindPath0(t *testing.T) {
	g := [][]int{
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 0, 0, 1},
	}

	start := Position{0, 0}
	end := Position{0, 3}

	actual := FindPath(start, end, g)
	expected := [][]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}, {1, 2}, {0, 2}, {0, 3}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Path does not match expected result\nExpected:\n%v\nReceived:\n%v", expected, actual)
	}
}

func TestFindPath1(t *testing.T) {
	g := [][]int{
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	}

	start := Position{0, 0}
	end := Position{11, 6}

	actual := FindPath(start, end, g)
	expected := [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {5, 1}, {5, 2}, {4, 2}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}, {7, 6}, {7, 7}, {7, 8}, {7, 9}, {7, 10}, {7, 11}, {8, 11}, {9, 11}, {9, 10}, {9, 9}, {9, 8}, {9, 7}, {10, 7}, {10, 6}, {11, 6}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Path does not match expected result\nExpected:\n%v\nReceived:\n%v", expected, actual)
	}
}

func TestFindPathNoPathExists(t *testing.T) {
	g := [][]int{
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 1},
	}

	start := Position{0, 0}
	end := Position{0, 3}

	actual := FindPath(start, end, g)

	if len(actual) != 0 {
		t.Errorf("Expected empty path")
	}
}

// TODO: Add out of bounds error / test
