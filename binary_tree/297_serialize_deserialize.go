package binary_tree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {

}

// Serializes a tree to a single string.
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	nodes = nil
	serialize(root)
	return strings.Join(nodes, ",")
}

var nodes []string

func serialize(root *TreeNode) {
	if root == nil {
		nodes = append(nodes, "0")
		return
	}

	nodes = append(nodes, strconv.Itoa(root.Val))
	serialize(root.Left)
	serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	deserializeArr := strings.Split(data, ",")
	if len(deserializeArr) == 0 {
		return nil
	}
	return deserialize(deserializeArr)
}

//
func deserialize(data []string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	val, _ := strconv.Atoi(data[0])
	if val == 0 {
		return nil
	}

	// 取第一个值作为根
	tree := &TreeNode{
		Val:   val,
		Left:  deserialize(data[1:]),
		Right: deserialize(data[1:]),
	}

	return tree
}
