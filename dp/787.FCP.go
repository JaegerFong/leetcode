package dp

var source int
var indegree map[int][]pNodeInfo

const MAX_VALUE = 9999999

type pNodeInfo struct {
	nidx   int
	weight int
}

// K站中转内最便宜的航班
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	source = src
	// 初始化节点的入度信息

	return 0
}

func fcpdp(dst, k int) int {
	if dst == source {
		return 0
	}

	if k <= 0 {
		return -1
	}

	res := MAX_VALUE
	if r, ok := indegree[dst]; ok {
		// 存在节点
		for _, v := range r {
			from := v.nidx
			weight := v.weight
			res = min(res, fcpdp(from, k-1)+weight)
		}
	}

	if res == MAX_VALUE {
		return -1
	}

	return res
}

// 问题转换:
// 1.有向图最小加权路径和,限制边数K
// 2.动态规划

// dp定义: dp(dst,k) = x , dst表示终点,k表示限制次数,x表示最少花费
// 前置条件: 记录一个节点的入度,也就是有多少个节点指向了该节点
// dp(dst,k) = min[dp(s1,k-1)+w1,dp(s2,k-1)+w2...dp(sn,k-1)+w3]
// 含义: 从所有指向dst节点中，选择最少花费
// s1,s2...sn: 指向dst的节点
// k-1: 上个节点的限制k
// w1,w2...wn: dst到sn节点的权重
