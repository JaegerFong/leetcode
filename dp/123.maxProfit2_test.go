package dp

import "testing"

func Test_maxProfit2(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	r := maxProfit2(prices)
	t.Log(r)
}
