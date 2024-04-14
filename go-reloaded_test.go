package main

import "testing"

type addTest struct {
	inputFile  string
	outputFile string
}

var addTests = []addTest{
	{"sample.txt", "Punctuation tests are... kinda boring, don't you think!?"},
}

func TestGoReloaded(t *testing.T) {
	for _, test := range addTests {
		if output := modify(test.inputFile); output != test.outputFile {
			t.Errorf("got %q, wanted %q", output, test.outputFile)
		}
	}
}
