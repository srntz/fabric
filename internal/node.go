package internal

type Node interface {
	Len() int
	Val() string
	ByteAt(i int) (byte, error)
	SplitAt(i int) (Node, Node, error)
}
