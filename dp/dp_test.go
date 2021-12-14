package dp

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	//fmt.Println(coinChange2(coins, amount))

	fmt.Println(coinChangeIteration(coins, amount))

}

func TestLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func TestFib(t *testing.T) {
	fmt.Println(fib_memo(5))
	fmt.Println(fib_dp_table(6))
}

func TestCoinchange(t *testing.T) {
	c := []int{1, 2, 5}
	t.Log(coinchange(c, 11))
}

func TestLis(t *testing.T) {
	arr := []int{1, 4, 3, 4, 2, 3}
	t.Log(lis(arr))
}

func TestPack01(t *testing.T) {
	w := 3
	n := 4
	wt := []int{2, 1, 3, 4}
	val := []int{4, 2, 3, 4}
	t.Log(pkdp(w, n, wt, val))
}

func TestVector(t *testing.T) {
	n := 3
	sum := 4
	dp := make([][]bool, n)

	// 初始化base case
	for i := 0; i < n; i++ {
		item := make([]bool, sum)
		dp[i] = append(dp[i], item...)
		for j := 0; j < sum; j++ {
			// weight==0,true
			if j == 0 {
				dp[i][j] = true
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < sum; j++ {
			t.Log(dp[i][j])
		}
	}
}
