package main

import (
	"testing"
)

// struct with numbers and expected fields
type data struct {
	numbers  string
	expected float64
}

// all test data
var testData = []data{
	{"11 22 33 44 55", 33},
	{"1 2 3 4", 2.5},
	{"1.11 2.22 3.33", 2.22},
	{"0.11 0.22", 0.17},
	{"0", 0},
	{"-1", -1},
	{"-0.11 -0.22", -0.17},
	{"-1.11 -2.22 -3.33", -2.22},
	{"-1 -2 -3 -4", -2.5},
	{"-11 -22 -33 -44 -55", -33},
}

// unit test for calAvgNum function
func TestCalAvgNum(t *testing.T) {
	for _, test := range testData {
		if actual := calAvgNum(test.numbers); actual != test.expected {
			t.Error("data =", test.numbers, "|", "actual =", actual, "|", "expected =", test.expected)
			t.Fail()
		}
	}
}
