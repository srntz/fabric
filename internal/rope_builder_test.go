package internal

import (
	"testing"

	"github.com/srntz/fabric/internal/spec"
)

func TestBuild(t *testing.T) {
	s := spec.RandomString(10000)
	rb := NewRopeBuilder(s)
	rope := rb.Build()

	ropelen := rope.root.Len()
	if ropelen != 10000 {
		t.Errorf("Unexpected Rope length. Want: %d, Got: %d", 10000, ropelen)
	}

	ropeval := rope.root.Val()
	if ropeval != s {
		t.Errorf("Unexpected Rope content. Want: %s, Got: %s", s, ropeval)
	}
}

func TestBlockifyString(t *testing.T) {
	s := spec.RandomString(10000)
	rb := RopeBuilder{s: s}
	rb.blockifyString()

	for i := range rb.blocks {
		if i == len(rb.blocks)-1 {
			continue
		}

		if len(rb.blocks[i]) != MAX_LEAF_CONTENT_LEN {
			t.Errorf("Unexpected block length. Want: %d, Got: %d", MAX_LEAF_CONTENT_LEN, len(rb.blocks[i]))
		}
	}
}
