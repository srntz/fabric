package internal

type Rope struct {
	root Node
}

func (r *Rope) ByteAt(index int) (byte, error) {
	return r.root.ByteAt(index)
}
