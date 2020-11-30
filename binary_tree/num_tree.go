package binary_tree

//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
/*
二叉搜索树的特性:
1. 左节点 < 节点 && 右节点 > 节点
2. 左子树 < 根节点 && 右子树 > 根节点

当左边数为0，右边为 n-1;
当左边数为1，右边为n-2;
....
当左边数为n-1，右边为0
即 G（n） = G(0)*G(n-1) + G(1)*G(n-2) + ... + G(n-1)*G(0)
*/
func numTrees(n int) int {
	count := 1
	for i := 0;i<n;i++ {
		count = (count*2*(2*i+1)) / (i+2)
	}
	return count
}
