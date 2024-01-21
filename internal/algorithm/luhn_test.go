package algorithm

import "testing"

type sumOfDigitsTest struct {
	given, expected int
}

var sumOfDigitsTests = []sumOfDigitsTest{
	{16, 7},
	{5, 5},
	{32, 5},
	{192, 12},
	{90231, 15},
}

func TestSumOfDigits(t *testing.T) {
	for _, test := range sumOfDigitsTests {
		t.Logf("Running %d, %d\n", test.given, test.expected)
		if output := sumOfDigits(test.given); output != test.expected {
			t.Errorf("Given %d, expected %d, got %d.\n", test.given, test.expected, output)
		}
	}
}

func TestReverseCardNumber(t *testing.T) {
	output := reverseCardNumber("123456")
	expected := []rune("654321")
	t.Logf("Expected %d, ouput %d\n", expected, output)
	if len(output) != len(expected) {
		t.Errorf("Mismatched length.")
	}
	for position, r := range output {
		if r != expected[position] {
			t.Errorf("Mismatched character at %d, expected %d, got %d\n", position, expected[position], r)
		}
	}
}

type removeLastCharacterTest struct {
	given, expected string
}

var removeLastCharacterTests = []removeLastCharacterTest{
	{"", ""},
	{"a", ""},
	{"abcef", "abce"},
}

func TestRemoveLastCharacter(t *testing.T) {
	for _, test := range removeLastCharacterTests {
		t.Logf("Running %s, %s\n", test.given, test.expected)
		if output := removeLastCharacter(test.given); output != test.expected {
			t.Errorf("Given %s, expected %s, got %s.\n", test.given, test.expected, output)
		}
	}
}
