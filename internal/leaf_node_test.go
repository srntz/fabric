package internal

import (
	"errors"
	"testing"
)

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

func TestLeafNodeSplitAt(t *testing.T) {
	tables := []struct {
		name                 string
		index                int
		nodeContent          string
		expectedLeftContent  string
		expectedRightContent string
		expectedError        error
	}{
		{
			name:                 "Split Leaf Node at negative index",
			index:                -1,
			nodeContent:          "my_node_content",
			expectedLeftContent:  "",
			expectedRightContent: "",
			expectedError:        IndexOutOfBoundsError,
		},
		{
			name:                 "Split Leaf Node at index greater then the bounds of the content by 1",
			index:                16,
			nodeContent:          "my_node_content",
			expectedLeftContent:  "",
			expectedRightContent: "",
			expectedError:        IndexOutOfBoundsError,
		},
		{
			name:                 "Split Leaf Node at index 0",
			index:                0,
			nodeContent:          "my_node_content",
			expectedLeftContent:  "",
			expectedRightContent: "my_node_content",
			expectedError:        nil,
		},
		{
			name:                 "Split Leaf Node at the index that equals the length of the content",
			index:                15,
			nodeContent:          "my_node_content",
			expectedLeftContent:  "my_node_content",
			expectedRightContent: "",
			expectedError:        nil,
		},
		{
			name:                 "Split Leaf Node at an index between the beginning and the end of the content",
			index:                9,
			nodeContent:          "my_node_content",
			expectedLeftContent:  "my_node_c",
			expectedRightContent: "ontent",
			expectedError:        nil,
		},
	}

	for _, row := range tables {
		t.Run(row.name, func(t *testing.T) {
			leaf := &LeafNode{content: row.nodeContent}
			left, right, err := leaf.SplitAt(row.index)

			if row.expectedError != nil {
				if err == nil {
					t.Fatalf("Unexpected error. Expected %s but got nil", row.expectedError.Error())
				}

				if !errors.Is(err, row.expectedError) {
					t.Errorf("Unexpected error. Want: %s, Got: %s", row.expectedError, err)
				}

				if left != nil || right != nil {
					t.Errorf("Unexpected node on error. Left and Right nodes must be nil. Got: Left - %s, Right - %s", left, right)
				}

				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: Expected nil but got %s", err.Error())
			}

			if left == nil || right == nil {
				t.Fatalf("Unexpected nil node in error-free context")
			}

			if left.Val() != row.expectedLeftContent {
				t.Errorf("Unexpected left content. Want: %s, Got: %s", row.expectedLeftContent, left.Val())
			}

			if right.Val() != row.expectedRightContent {
				t.Errorf("Unexpected right content. Want: %s, Got: %s", row.expectedRightContent, right.Val())
			}
		})
	}
}
