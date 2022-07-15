package logical_or_of_two_binary_grids_represented_as_quad_trees

type Node struct {
	Val, IsLeaf                                bool
	TopLeft, TopRight, BottomLeft, BottomRight *Node
}

func Intersect(t1, t2 *Node) *Node {
	if t1 == nil || t2 == nil {
		return nil
	}
	if t1.IsLeaf {
		if t1.Val {
			return t1
		}
		return t2
	}
	if t2.IsLeaf {
		if t2.Val {
			return t2
		}
		return t1
	}
	topLeft := Intersect(t1.TopLeft, t2.TopLeft)
	topRight := Intersect(t1.TopRight, t2.TopRight)
	bottomLeft := Intersect(t1.BottomLeft, t2.BottomLeft)
	bottomRight := Intersect(t1.BottomRight, t2.BottomRight)
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &Node{
			Val:    topLeft.Val,
			IsLeaf: true,
		}
	}
	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
}
