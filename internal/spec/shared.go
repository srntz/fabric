package spec

import "math/rand"

func RandomString(len int) string {
	buf := make([]byte, len)
	for i := range len {
		buf[i] = byte(rand.Intn(10) + '0')
	}
	return string(buf)
}
