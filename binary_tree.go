package myplayground

// Node represents a Node in the tree. It contains generic data, a left node
// and a right node.
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

// DFS performs a Depth first traversal of the tree and return the list of
// interface{} data. Uses a recursive algorithm for traversal.
func (n *Node) DFS() []interface{} {
	if n == nil {
		// TODO: should this return an empty list?
		return nil
	}

	ret := []interface{}{}
	ret = append(ret, n.Data)
	if n.Left != nil {
		for _, element := range n.Left.DFS() {
			ret = append(ret, element)
		}
	}
	if n.Right != nil {
		for _, element := range n.Right.DFS() {
			ret = append(ret, element)
		}
	}

	return ret
}

// BFS performs a Breadth first traversal of the tree and return the list of
// interface{} data. Uses a queue data structure for traversal.
func (n *Node) BFS() []interface{} {
	if n == nil {
		// TODO: should this return an empty list?
		return nil
	}

	ret := []interface{}{}
	q := []*Node{n}
	for len(q) != 0 {
		current := q[0]
		q = q[1:len(q)]
		ret = append(ret, current.Data)

		if current.Left != nil {
			q = append(q, current.Left)
		}
		if current.Right != nil {
			q = append(q, current.Right)
		}

	}

	return ret
}

// Size returns the amount of nodes in the tree.
func (n *Node) Size() int {
	if n == nil {
		return 0
	}

	size := 1
	size += n.Left.Size()
	size += n.Right.Size()

	return size
}

// Height returns the height (or number of levels) of the tree.
func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	height := 1
	lh := n.Left.Height()
	rh := n.Right.Height()

	max := lh
	if rh > max {
		max = rh
	}
	height += max

	return height
}
