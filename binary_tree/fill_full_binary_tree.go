package binary_tree

// 填充二叉树节点的右侧指针 每个节点的下一个右侧节点指针,即将同一层级的节点串起来


func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	ConnectTwo(root.Left,root.Right)
	return root
}

func ConnectTwo(node1 *Node,node2 *Node){
	if node1 == nil || node2==nil{
		return
	}

	node1.Next = node2

	// 分别链接各自子树
	ConnectTwo(node1.Left,node1.Right)
	ConnectTwo(node2.Left,node2.Right)

	// 链接两个不直接相连的节点
	ConnectTwo(node1.Right,node2.Left)
}