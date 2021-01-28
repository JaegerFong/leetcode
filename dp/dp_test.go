package dp

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange2(coins, amount))
}

func TestDpFib2(t *testing.T) {
	fmt.Println(dp_fib2(5))
}

func TestLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
