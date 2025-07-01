package internal

type BranchNode struct {
	left   Node
	right  Node
	weight int
}

func (bn *BranchNode) Len() int {
	return bn.weight + bn.right.Len()
}

func (bn *BranchNode) Val() string {
	return bn.right.Val() + bn.left.Val()
}
