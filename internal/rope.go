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
	left, right, err := r.root.SplitAt(i)
	if err != nil {
		return nil, nil, err
	}

	return NewRope(left), NewRope(right), nil
}
