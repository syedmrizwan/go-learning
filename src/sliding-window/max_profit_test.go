package sliding_window

import "testing"

func TestMaxProfit(t *testing.T) {
	subTests := []struct {
		input  []int
		result int
	}{
		{
			input:  []int{7, 1, 5, 3, 6, 4},
			result: 5,
		},
	}

	for _, test := range subTests {
		if s := maxProfit(test.input); s != test.result {
			t.Errorf("wanted %v, got %v", test.result, s)
		}
	}
}
