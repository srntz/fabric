package internal

import "testing"

func TestLeafNodeByteAt(t *testing.T) {
	l := LeafNode{
		content: "abcdefg",
	}

	b, err := l.ByteAt(5)
	if err != nil {
		t.Error(err.Error())
	}

	if b != l.content[5] {
		t.Errorf("Unexpected byte at index. Want %b, Got: %b", l.content[5], b)
	}
}
