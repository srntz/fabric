package fabric

import "github.com/srntz/fabric/internal"

func BuildRopeFromString(s string) *internal.Rope {
	return internal.NewRopeBuilder(s).Build()
}
