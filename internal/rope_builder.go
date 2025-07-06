package internal

type RopeBuilder struct {
	s      string
	blocks []string
}

func NewRopeBuilder(s string) *RopeBuilder {
	rb := &RopeBuilder{s: s}
	rb.blockifyString()
	return rb
}

func (rb *RopeBuilder) Build() *Rope {
	return NewRope(build(rb.blocks))
}

func build(blocks []string) Node {
	if len(blocks) == 1 {
		return &LeafNode{content: blocks[0]}
	}

	mid := len(blocks) / 2
	left := build(blocks[:mid])
	right := build(blocks[mid:])
	return &BranchNode{
		left:   left,
		right:  right,
		weight: left.Len(),
	}
}

func (rb *RopeBuilder) blockifyString() {
	rb.blocks = []string{}
	for i := 0; i < len(rb.s); i += MAX_LEAF_CONTENT_LEN {
		if i+MAX_LEAF_CONTENT_LEN > len(rb.s) {
			rb.blocks = append(rb.blocks, rb.s[i:len(rb.s)])
		} else {
			rb.blocks = append(rb.blocks, rb.s[i:i+MAX_LEAF_CONTENT_LEN])
		}
	}
}
