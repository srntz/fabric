package internal

import (
	"testing"

	"github.com/srntz/fabric/internal/spec"
)

func TestRopeByteAt(t *testing.T) {
	s := spec.RandomString(10000)
	rope := NewRopeBuilder(s).Build()

	byterope, err := rope.ByteAt(7979)
	if err != nil {
		t.Error(err.Error())
	}

	bytestr := s[7979]
	if byterope != bytestr {
		t.Errorf("Unexpected byte at index. Want: %b, Got: %b", bytestr, byterope)
	}
}
