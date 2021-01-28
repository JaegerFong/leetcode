package dp

func dp_fib(n int) int {
	arr := make([]int, n)
	arr[1], arr[2] = 1, 1
	for i := 3; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

func dp_fib2(n int) int {
	if n == 2 || n == 1 {
		return 1
	}
	prev, cur := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + cur
		prev = cur
		cur = sum
	}

	return cur
}
