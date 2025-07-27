package internal

import (
	"errors"
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

func TestRopeSplitAt(t *testing.T) {
	content := spec.RandomString(10000)
	table := []struct {
		name                 string
		initialContent       string
		index                int
		expectedLeftContent  string
		expectedRightContent string
		expectedError        error
	}{
		{
			name:           "Split rope at negative index",
			initialContent: content,
			index:          -1,
			expectedError:  IndexOutOfBoundsError,
		},
		{
			name:           "Split rope at index greater than rope length",
			initialContent: content,
			index:          len(content) + 1,
			expectedError:  IndexOutOfBoundsError,
		},
		{
			name:                 "Split rope at index 0",
			initialContent:       content,
			index:                0,
			expectedLeftContent:  "",
			expectedRightContent: content,
		},
		{
			name:                 "Split rope at last index",
			initialContent:       content,
			index:                len(content),
			expectedLeftContent:  content,
			expectedRightContent: "",
		},
		{
			name:                 "Split rope near the middle",
			initialContent:       content,
			index:                6279,
			expectedLeftContent:  content[:6279],
			expectedRightContent: content[6279:],
		},
	}

	for _, row := range table {
		t.Run(row.name, func(t *testing.T) {
			rope := NewRopeBuilder(row.initialContent).Build()
			left, right, err := rope.SplitAt(row.index)

			if row.expectedError != nil {
				if err == nil {
					t.Fatalf("Expected error: %s. Got nil.", row.expectedError)
				}

				if !errors.Is(err, row.expectedError) {
					t.Fatalf("Expected error: %s. Got: %s.", row.expectedError, err)
				}

				if left != nil || right != nil {
					t.Errorf("Unexpected rope in return. Expected both left and right to be nil but got non-nil values.")
				}

				return
			}

			if err != nil {
				t.Fatalf("Received error in error-free context. Got: %s.", err)
			}

			if left == nil || right == nil {
				t.Fatalf("Unexpected rope in return. Expected Rope but got nil values.")
			}

			leftValue := left.root.Val()
			if leftValue != row.expectedLeftContent {
				t.Errorf("Unexpected content in left rope. Want: %s. Got: %s.", row.expectedLeftContent, leftValue)
			}

			rightValue := right.root.Val()
			if rightValue != row.expectedRightContent {
				t.Errorf("Unexpected content in right rope. Want: %s. Got: %s.", row.expectedRightContent, rightValue)
			}
		})
	}
}

func TestRopeConcat(t *testing.T) {
	table := []struct {
		name             string
		leftRopeContent  string
		rightRopeContent string
	}{
		{
			name:             "Concatenate to an empty rope",
			leftRopeContent:  "",
			rightRopeContent: "to concat",
		},
		{
			name:             "Concatenate an empty rope",
			leftRopeContent:  "original rope",
			rightRopeContent: "",
		},
		{
			name:             "Concatenate two non-empty ropes",
			leftRopeContent:  "original rope",
			rightRopeContent: "to concat",
		},
	}

	for _, row := range table {
		leftRope := NewRopeBuilder(row.leftRopeContent).Build()
		rightRope := NewRopeBuilder(row.rightRopeContent).Build()

		t.Run(row.name, func(t *testing.T) {
			newRope := Concat(leftRope, rightRope)

			want := row.leftRopeContent + row.rightRopeContent
			got := newRope.root.Val()
			if got != want {
				t.Fatalf("Unexpected concatenated value. Want: %s. Got: %s", want, got)
			}
		})
	}
}

func TestRopeInsertAtWithValidInput(t *testing.T) {
	table := []struct {
		name                string
		originalRopeContent string
		contentToInsert     string
		insertAt            int
		expect              string
	}{
		{
			name:                "inserting at index 0",
			originalRopeContent: "rope content",
			contentToInsert:     "inserted ",
			insertAt:            0,
			expect:              "inserted rope content",
		},
		{
			name:                "inserting at the last index",
			originalRopeContent: "rope content",
			contentToInsert:     " inserted",
			insertAt:            12,
			expect:              "rope content inserted",
		},
		{
			name:                "inserting in the middle of the rope",
			originalRopeContent: "rope content",
			contentToInsert:     "inserted",
			insertAt:            2,
			expect:              "roinsertedpe content",
		},
	}

	for _, row := range table {
		rope := NewRopeBuilder(row.originalRopeContent).Build()

		t.Run(row.name, func(t *testing.T) {
			newRope, err := rope.InsertAt(row.insertAt, row.contentToInsert)
			if err != nil {
				t.Fatalf("Unexpected Error: %s", err.Error())
			}

			want := row.expect
			got := newRope.root.Val()
			if got != want {
				t.Fatalf("Unexpected inserted content. Want: %s, Got: %s", want, got)
			}
		})
	}
}

func TestRopeInsertAtWithInvalidInput(t *testing.T) {
	table := []struct {
		name                string
		originalRopeContent string
		contentToInsert     string
		insertAt            int
	}{
		{
			name:                "inserting at a negative index",
			originalRopeContent: "rope content",
			contentToInsert:     "inserted",
			insertAt:            -1,
		},
		{
			name:                "inserting at an index greater than the size of the original content",
			originalRopeContent: "rope content",
			contentToInsert:     "inserted",
			insertAt:            13,
		},
	}

	for _, row := range table {
		rope := NewRopeBuilder(row.originalRopeContent).Build()
		t.Run(row.name, func(t *testing.T) {
			newRope, err := rope.InsertAt(row.insertAt, row.contentToInsert)
			if err == nil {
				t.Fatal("Expected error but got nil")
			}

			if newRope != nil {
				t.Fatalf("Expected nil rope in an error-returning context. Got: %s", newRope)
			}
		})
	}
}
