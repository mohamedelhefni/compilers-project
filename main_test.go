package main

import (
	"fmt"
	"testing"
)

type ParseTestCase struct {
	Input  string
	ShouldFail bool
}

func TestParser(t *testing.T) {
	testCases := []ParseTestCase{
		{Input: "x = 5", ShouldFail: false},
		{Input: "if (x > 2) { y = 4 } else { y = 5 }", ShouldFail: false},
		{Input: "x = (3 + 2) - (2 * 1)", ShouldFail: false},
		{Input: "if (x > 2) { 3 = 4 } else { y = 5 }", ShouldFail: true}, 
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			lexer := NewLexer(testCase.Input)
			tokens := lexer.tokenize()

			parser := NewParser(tokens)
			err := parser.parse()

			if (err != nil) != testCase.ShouldFail {
				t.Errorf("Test failed for input: %s, expected failure: %t", testCase.Input, testCase.ShouldFail)
			}
		})
	}
}

