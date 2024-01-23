package common

import "testing"

type getNthDigitTest struct {
	given    string
	n        int
	expected int
}

var getNthDigitTests = []getNthDigitTest{
	{"123456789", 0, 1},
	{"123456789", 1, 2},
	{"123456789", 8, 9},
	{"1", 0, 1},
}

func TestGetNthDigit(t *testing.T) {
	for _, test := range getNthDigitTests {
		t.Logf("Running %s, %d\n", test.given, test.expected)
		if output, _ := getNthDigit(test.given, test.n); output != test.expected {
			t.Errorf("Given %s, expected %d, got %d.\n", test.given, test.expected, output)
		}
	}
}

func TestGetNthDigitError(t *testing.T) {
	if _, err := getNthDigit("123", 5); err == nil {
		t.Fail()
	}
	if _, err := getNthDigit("abcef", 1); err == nil {
		t.Fail()
	}

}
