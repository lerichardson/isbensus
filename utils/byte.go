package utils

import (
	"crypto/rand"
)

func RandomBytes(size int) (blk []byte) {
	blk = make([]byte, size)
	_, err := rand.Read(blk)
	if err != nil {
		panic(err)
	}
	return
}
