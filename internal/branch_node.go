package internal

import "strings"

type BranchNode struct {
	left   Node
	right  Node
	weight int
}

func (bn *BranchNode) Len() int {
	return bn.weight + bn.right.Len()
}

func (bn *BranchNode) Val() string {
	b := strings.Builder{}
	b.WriteString(bn.left.Val())
	b.WriteString(bn.right.Val())
	return b.String()
}

func (bn *BranchNode) ByteAt(i int) (byte, error) {
	if i > bn.weight {
		return bn.right.ByteAt(i - bn.weight)
	}
	return bn.left.ByteAt(i)
}
