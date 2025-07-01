package internal

var MAX_LEAF_CONTENT_LEN int = 4096
var MIN_LEAF_CONTENT_LEN int = 2048

type LeafNode struct {
	content string
}

func (ln *LeafNode) Len() int {
	return len(ln.content)
}

func (ln *LeafNode) Val() string {
	return ln.content
}
