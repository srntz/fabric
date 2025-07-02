package fabric

import "github.com/srntz/fabric/internal"

func BuildRopeFromString(s string) *FabricRope {
	return &FabricRope{
		rope: internal.NewRopeBuilder(s).Build(),
	}
}
