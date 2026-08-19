package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
}{
	{
		input: " hello world ",
		expected: []string{"hello", "world"},
	},

	{
		input: "TeStInG LOWERCASE",
		expected: []string{"testing", "lowercase"},
	},

	{	input: "ONEword",
		expected: []string{"oneword"},
	},

	{	input: "",
		expected: []string{},
	},


}

for _,c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("expected: %v, got: %v", c.expected, actual)
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("expected word: %s, got: %s", expectedWord, word)

	}
}
}
}
