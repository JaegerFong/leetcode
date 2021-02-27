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

func TestLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func TestFib(t *testing.T) {
	fmt.Println(fib_memo(5))
	fmt.Println(fib_dp_table(6))
}
