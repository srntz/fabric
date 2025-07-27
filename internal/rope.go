package internal

type Rope struct {
	root Node
}

func NewRope(root Node) *Rope {
	return &Rope{root: root}
}

func (r *Rope) ByteAt(i int) (byte, error) {
	return r.root.ByteAt(i)
}

func (r *Rope) SplitAt(i int) (*Rope, *Rope, error) {
	if i > r.root.Len() || i < 0 {
		return nil, nil, IndexOutOfBoundsError
	}

	left, right, err := r.root.SplitAt(i)
	if err != nil {
		return nil, nil, err
	}

	return NewRope(left), NewRope(right), nil
}

func Concat(left *Rope, right *Rope) *Rope {
	return NewRope(
		&BranchNode{
			left:   left.root,
			right:  right.root,
			weight: left.root.Len(),
		},
	)
}

func (r *Rope) InsertAt(i int, seq string) (*Rope, error) {
	leftSplit, rightSplit, err := r.SplitAt(i)
	if err != nil {
		return nil, err
	}

	seqRope := NewRopeBuilder(seq).Build()
	return Concat(Concat(leftSplit, seqRope), rightSplit), nil
}
