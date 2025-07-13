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

func (bn *BranchNode) SplitAt(i int) (Node, Node, error) {
	if i > bn.weight {
		leftSplit, rightSplit, err := bn.right.SplitAt(i - bn.weight)
		if err != nil {
			return nil, nil, err
		}

		return &BranchNode{
			left:   bn.left,
			right:  leftSplit,
			weight: bn.weight,
		}, rightSplit, nil
	} else {
		leftSplit, rightSplit, err := bn.left.SplitAt(i)
		if err != nil {
			return nil, nil, err
		}

		return leftSplit, &BranchNode{
			left:   rightSplit,
			right:  bn.right,
			weight: rightSplit.Len(),
		}, nil
	}
}
