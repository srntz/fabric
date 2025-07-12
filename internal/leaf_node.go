package internal

import "errors"

var MAX_LEAF_CONTENT_LEN int = 2048
var MIN_LEAF_CONTENT_LEN int = 512

var IndexOutOfBoundsError = errors.New("Error: Index out of bounds")

type LeafNode struct {
	content string
}

func (ln *LeafNode) Len() int {
	return len(ln.content)
}

func (ln *LeafNode) Val() string {
	return ln.content
}

func (ln *LeafNode) ByteAt(i int) (byte, error) {
	if i > ln.Len() {
		return ' ', IndexOutOfBoundsError
	}
	return ln.content[i], nil
}

func (ln *LeafNode) SplitAt(i int) (Node, Node, error) {
	if i > ln.Len() || i < 0 {
		return nil, nil, IndexOutOfBoundsError
	}
	return &LeafNode{content: ln.content[:i]}, &LeafNode{content: ln.content[i:]}, nil
}
