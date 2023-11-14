package common

import (
	"math/rand"
)

const dicts = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func alphanum(n int) string {
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = dicts[rand.Intn(len(dicts))]
	}
	return string(bs)
}
