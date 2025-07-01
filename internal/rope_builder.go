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
	return &Rope{}
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
