package stack

import "testing"

func TestIsValid(t *testing.T) {
	subTests := []struct {
		input  string
		result bool
	}{
		{
			input:  "()",
			result: true,
		}, {
			input:  "()[]{}",
			result: true,
		},
		{
			input:  "(]",
			result: false,
		},
	}

	for _, test := range subTests {
		if s := isValid(test.input); s != test.result {
			t.Errorf("wanted %t, got %t", test.result, s)
		}
	}
}
