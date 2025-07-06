package internal

import "errors"

var MAX_LEAF_CONTENT_LEN int = 2048
var MIN_LEAF_CONTENT_LEN int = 512

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
		return ' ', errors.New("index out of bounds")
	}
	return ln.content[i], nil
}

func (ln *LeafNode) SplitAt(i int) (Node, Node, error) {
	if i > ln.Len() {
		return nil, nil, errors.New("index out of bounds")
	}
	return &LeafNode{content: ln.content[:ln.Len()/2]}, &LeafNode{content: ln.content[ln.Len()/2:]}, nil
}
