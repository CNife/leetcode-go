package serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"

	"github.com/CNife/leetcode-go/types"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) Serialize(tree *types.TreeNode) string {
	var result []string
	queue := []*types.TreeNode{tree}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, "")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}
	for len(result) > 0 && len(result[len(result)-1]) == 0 {
		result = result[:len(result)-1]
	}
	return strings.Join(result, ",")
}

func (c *Codec) Deserialize(data string) *types.TreeNode {
	nodes := strings.Split(data, ",")
	if len(nodes[0]) == 0 {
		return nil
	}

	root := parseNode(nodes[0])
	queue := []*types.TreeNode{root}
	for i := 1; i < len(nodes); {
		node := queue[0]
		queue = queue[1:]

		node.Left = parseNode(nodes[i])
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nodes) {
			node.Right = parseNode(nodes[i])
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

func parseNode(value string) *types.TreeNode {
	if len(value) == 0 {
		return nil
	} else {
		nodeVal, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return &types.TreeNode{Val: nodeVal}
	}
}
