package binary_search

import "testing"

func TestSearch(t *testing.T) {
	subTests := []struct {
		input  []int
		target int
		result int
	}{
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			result: 4,
		},
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			result: -1,
		},
	}

	for _, test := range subTests {
		if s := search(test.input, test.target); s != test.result {
			t.Errorf("wanted %v, got %v", test.result, s)
		}
	}
}
