package main

import (
	"testing"
)

func TestRecursivelyMemoized(t *testing.T) {
	t.Run("small-num", testFib(6, 8))
	t.Run("medium-num", testFib(50, 12586269025))
}

func testFib(input int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		got := RecursivelyMemoized(input)
		if got != expected {
			t.Errorf("Incorrect Fibonacci calculation, got: %d, want %d.", got, expected)
		}
	}
}

func TestConvertInput(t *testing.T) {
	t.Run("zero", testConvertInput("0", 0))
	t.Run("trivial-num", testConvertInput("5", 5))
	t.Run("non-num", testConvertInput("this is a str", -1))
	t.Run("negative", testConvertInput("-100", -1))
}

func testConvertInput(input string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		got := convertInput(input)
		if got != expected {
			t.Errorf("Input conversion incorrect, got: %d, want %d.", got, expected)
		}
	}
}
