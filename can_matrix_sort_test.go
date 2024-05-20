package main

import (
	"fmt"
	"strings"
	"testing"
)

const dimensionBigMatrix = 100

var firstLineBigMatrix = fmt.Sprintf("%d\n", dimensionBigMatrix) //nolint:gochecknoglobals

func generateBigMatrix() string {
	line := "1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1" +
		" 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1\n"

	result := make([]byte, len(firstLineBigMatrix)+len(line)*100)

	copy(result, firstLineBigMatrix)

	for i := 0; i < dimensionBigMatrix; i++ {
		copy(result[i*len(line)+len(firstLineBigMatrix):], line)
	}

	return string(result)
}

type testCaseCanMatrixSort struct {
	name           string
	input          string
	expectedOutput bool
}

var testCases = [...]testCaseCanMatrixSort{ //nolint:gochecknoglobals
	{
		name: "Example 1 from task.md",
		input: `2
	1 2
	2 1`,
		expectedOutput: true,
	},
	{
		name: "Example 2 from task.md",
		input: `3
	10 20 30
	1 1 1
	0 0 1`,
		expectedOutput: false,
	},
	{
		name: "big values",
		input: `2
	1000000000 1000000002
	1000000002 1000000000`,
		expectedOutput: true,
	},
	{
		name: "zero values",
		input: `3
	0 0 0
	0 0 0
	0 0 0`,
		expectedOutput: true,
	},
	{
		name: "empty containers",
		input: `3
	0 0 1
	0 0 0
	0 0 0`,
		expectedOutput: true,
	},
	{
		name: "can not sort containers",
		input: `3
	0 0 3
	0 3 0
	1 0 1`,
		expectedOutput: false,
	},
	{
		name: "dimension matrix = 1",
		input: `1
	4`,
		expectedOutput: true,
	},
	{
		name:           "big matix",
		input:          generateBigMatrix(),
		expectedOutput: true,
	},
}

func TestCanMatrixSortWithCycles(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result, err := CanMatrixSortWithCycles(strings.NewReader(testCase.input))
			if err != nil {
				t.Fatal(err)
			}

			if result != testCase.expectedOutput {
				t.Errorf("expected %v, got %v", testCase.expectedOutput, result)
			}
		})
	}
}

func TestCanMatrixSortWithMap(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result, err := CanMatrixSortWithMap(strings.NewReader(testCase.input))
			if err != nil {
				t.Fatal(err)
			}

			if result != testCase.expectedOutput {
				t.Errorf("expected %v, got %v", testCase.expectedOutput, result)
			}
		})
	}
}

func TestCanMatrixSortWithSort(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result, err := CanMatrixSortWithSort(strings.NewReader(testCase.input))
			if err != nil {
				t.Fatal(err)
			}

			if result != testCase.expectedOutput {
				t.Errorf("expected %v, got %v", testCase.expectedOutput, result)
			}
		})
	}
}
