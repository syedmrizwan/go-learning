package two_pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	subTests := []struct {
		input  string
		result bool
	}{
		{
			input:  "A man, a plan, a canal: Panama",
			result: true,
		}, {
			input:  "race a car",
			result: false,
		},
	}

	for _, test := range subTests {
		if s := isPalindrome(test.input); s != test.result {
			t.Errorf("wanted %t, got %t", test.result, s)
		}
	}
}
