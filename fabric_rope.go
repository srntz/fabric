package fabric

import "github.com/srntz/fabric/internal"

type FabricRope struct {
	rope *internal.Rope
}

func (fr *FabricRope) ByteAt(index int) {
	fr.rope.ByteAt(index)
}
