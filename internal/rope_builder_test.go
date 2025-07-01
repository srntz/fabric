package internal

import (
	"math/rand"
	"testing"
)

func TestBlockifyString(t *testing.T) {
	s := RandomString(10000)
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

func RandomString(len int) string {
	buf := make([]byte, len)
	for i := range len {
		buf[i] = byte(rand.Intn(10) + '0')
	}
	return string(buf)
}
