package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// add more cases here
	{
		input:    "Charmander Bulbasaur PIKACHU",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
}

func TestCleanInput(t *testing.T) {
	// ...
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of output not equal to expected")
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("The actual word %s not equal to expected word %s", actual[i], c.expected[i])
			}
		}
	}
	fmt.Printf("All tests passed")

}
